package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

// loadConfig must not touch clem.yaml for commands that don't need it.
func TestLoadConfigSkipsConfigForSpecialCommands(t *testing.T) {
	tmp := t.TempDir() // no clem.yaml here

	prev, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(prev) }()

	for _, name := range []string{"update", "init", "vault"} {
		cmd := &cobra.Command{Use: name}
		if err := loadConfig(cmd, nil); err != nil {
			t.Errorf("loadConfig(%q) returned error: %v", name, err)
		}
	}
}

func TestLoadConfigRequiresConfigForOtherCommands(t *testing.T) {
	tmp := t.TempDir() // no clem.yaml

	prev, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(prev) }()

	cmd := &cobra.Command{Use: "provision"}
	if err := loadConfig(cmd, nil); err == nil {
		t.Error("loadConfig(provision) expected error without clem.yaml, got nil")
	}
}
