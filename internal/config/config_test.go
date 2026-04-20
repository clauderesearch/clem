package config

import (
	"os"
	"path/filepath"
	"testing"
)

func writeYAML(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "clem.yaml")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("writing yaml: %v", err)
	}
	return path
}

func minYAML(caveman string) string {
	cavemanLine := ""
	if caveman != "" {
		cavemanLine = "\n    caveman: " + caveman
	}
	return `
project: myteam
coordination:
  backend: discord
  server_id: "1"
  channels: {general: "g"}
agents:
  lead:
    name: "Lead"
    model: "claude-sonnet-4-6"` + cavemanLine + "\n"
}

func TestCavemanLevel_StringLevels(t *testing.T) {
	cases := []struct {
		yaml    string
		want    CavemanLevel
		enabled bool
	}{
		{"lite", CavemanLite, true},
		{"full", CavemanFull, true},
		{"ultra", CavemanUltra, true},
		{"off", CavemanOff, false},
	}
	for _, tc := range cases {
		path := writeYAML(t, minYAML(tc.yaml))
		cfg, err := Load(path)
		if err != nil {
			t.Fatalf("caveman=%q: Load: %v", tc.yaml, err)
		}
		got := cfg.Agents["lead"].Caveman
		if got != tc.want {
			t.Errorf("caveman=%q: got %q, want %q", tc.yaml, got, tc.want)
		}
		if got.Enabled() != tc.enabled {
			t.Errorf("caveman=%q: Enabled()=%v, want %v", tc.yaml, got.Enabled(), tc.enabled)
		}
		if tc.enabled && got.Level() != tc.yaml {
			t.Errorf("caveman=%q: Level()=%q, want %q", tc.yaml, got.Level(), tc.yaml)
		}
	}
}

func TestCavemanLevel_LegacyBool(t *testing.T) {
	// true → ultra
	path := writeYAML(t, minYAML("true"))
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("caveman=true: Load: %v", err)
	}
	if got := cfg.Agents["lead"].Caveman; got != CavemanUltra {
		t.Errorf("caveman=true: got %q, want %q", got, CavemanUltra)
	}

	// false → off
	path = writeYAML(t, minYAML("false"))
	cfg, err = Load(path)
	if err != nil {
		t.Fatalf("caveman=false: Load: %v", err)
	}
	if got := cfg.Agents["lead"].Caveman; got != CavemanOff {
		t.Errorf("caveman=false: got %q, want %q", got, CavemanOff)
	}
}

func TestCavemanLevel_UnsetIsOff(t *testing.T) {
	path := writeYAML(t, minYAML(""))
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Agents["lead"].Caveman.Enabled() {
		t.Error("expected caveman disabled when unset")
	}
}

func TestCavemanLevel_InvalidStringRejectsAtLoad(t *testing.T) {
	path := writeYAML(t, minYAML("maximum"))
	_, err := Load(path)
	if err == nil {
		t.Fatal("expected error for invalid caveman value, got nil")
	}
}

func TestLoad_PrimaryMilestoneParsed(t *testing.T) {
	path := writeYAML(t, `
project: myteam
primary_milestone: "Ship v1 by 2027-01-01"
coordination:
  backend: discord
  server_id: "1"
  channels: {general: "g", tasks: "t"}
agents:
  lead:
    name: "Amara"
    role: "Lead"
    model: "claude-sonnet-4-6"
`)
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.PrimaryMilestone != "Ship v1 by 2027-01-01" {
		t.Errorf("PrimaryMilestone = %q", cfg.PrimaryMilestone)
	}
}

func TestLoad_PrimaryMilestoneOptional(t *testing.T) {
	path := writeYAML(t, `
project: myteam
coordination:
  backend: discord
  server_id: "1"
  channels: {general: "g"}
agents:
  lead:
    name: "Amara"
    model: "claude-sonnet-4-6"
`)
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.PrimaryMilestone != "" {
		t.Errorf("PrimaryMilestone = %q, want empty", cfg.PrimaryMilestone)
	}
}

func subagentYAML(subagentModel string) string {
	line := ""
	if subagentModel != "" {
		line = "\n    subagent_model: " + subagentModel
	}
	return `
project: myteam
coordination:
  backend: discord
  server_id: "1"
  channels: {general: "g"}
agents:
  lead:
    name: "Lead"
    model: "claude-opus-4-7"` + line + "\n"
}

func TestLoad_SubagentModelDefaultsWhenUnset(t *testing.T) {
	path := writeYAML(t, subagentYAML(""))
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got := cfg.Agents["lead"].SubagentModel; got != DefaultSubagentModel {
		t.Errorf("SubagentModel = %q, want %q", got, DefaultSubagentModel)
	}
}

func TestLoad_SubagentModelExplicitValuePreserved(t *testing.T) {
	path := writeYAML(t, subagentYAML(`"claude-haiku-4-5-20251001"`))
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got := cfg.Agents["lead"].SubagentModel; got != "claude-haiku-4-5-20251001" {
		t.Errorf("SubagentModel = %q, want %q", got, "claude-haiku-4-5-20251001")
	}
}

func TestLoad_SubagentModelOffDisables(t *testing.T) {
	path := writeYAML(t, subagentYAML("off"))
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got := cfg.Agents["lead"].SubagentModel; got != "" {
		t.Errorf("SubagentModel = %q, want empty (disabled)", got)
	}
}

func TestLoad_SubagentModelNoDefaultForOllama(t *testing.T) {
	yaml := `
project: myteam
coordination:
  backend: discord
  server_id: "1"
  channels: {general: "g"}
agents:
  lead:
    name: "Lead"
    model: "llama3"
    provider: ollama
`
	path := writeYAML(t, yaml)
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got := cfg.Agents["lead"].SubagentModel; got != "" {
		t.Errorf("SubagentModel = %q, want empty (ollama cannot use CLAUDE_CODE_SUBAGENT_MODEL)", got)
	}
}

func TestLoad_SubagentModelDefaultsForBedrock(t *testing.T) {
	yaml := `
project: myteam
coordination:
  backend: discord
  server_id: "1"
  channels: {general: "g"}
agents:
  lead:
    name: "Lead"
    model: "claude-opus-4-7"
    provider: bedrock
`
	path := writeYAML(t, yaml)
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got := cfg.Agents["lead"].SubagentModel; got != DefaultSubagentModel {
		t.Errorf("SubagentModel = %q, want %q (bedrock is Anthropic-backed)", got, DefaultSubagentModel)
	}
}

func TestLoad_GitIdentityParsed(t *testing.T) {
	path := writeYAML(t, `
project: myteam
coordination:
  backend: discord
  server_id: "1"
  channels: {general: "g"}
agents:
  lead:
    name: "Ada"
    model: "claude-sonnet-4-6"
    git_name: "clauderesearch"
    git_email: "212849679+clauderesearch@users.noreply.github.com"
`)
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	ac := cfg.Agents["lead"]
	if ac.GitName != "clauderesearch" {
		t.Errorf("GitName = %q, want %q", ac.GitName, "clauderesearch")
	}
	if ac.GitEmail != "212849679+clauderesearch@users.noreply.github.com" {
		t.Errorf("GitEmail = %q, want %q", ac.GitEmail, "212849679+clauderesearch@users.noreply.github.com")
	}
}

func TestLoad_GitIdentityOptional(t *testing.T) {
	path := writeYAML(t, minYAML(""))
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	ac := cfg.Agents["lead"]
	if ac.GitName != "" || ac.GitEmail != "" {
		t.Errorf("expected empty git identity when unset, got name=%q email=%q", ac.GitName, ac.GitEmail)
	}
}
