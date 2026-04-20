package cmd

import (
	"os"
	"strings"
	"testing"
)

func TestInitTemplateContainsRunnerExitProtocol(t *testing.T) {
	tmp := t.TempDir()

	prev, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(prev) }()

	if err := runInit(nil, nil); err != nil {
		t.Fatalf("runInit failed: %v", err)
	}

	data, err := os.ReadFile("CLAUDE.shared.md")
	if err != nil {
		t.Fatalf("reading CLAUDE.shared.md: %v", err)
	}
	content := string(data)

	checks := []string{
		"kill $PPID",
		"runner exit protocol",
		"How your session ends",
	}
	for _, want := range checks {
		if !strings.Contains(content, want) {
			t.Errorf("CLAUDE.shared.md missing %q", want)
		}
	}
}
