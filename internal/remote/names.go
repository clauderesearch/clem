package remote

import (
	"fmt"
	"regexp"
)

var validRepoName = regexp.MustCompile(`^[a-zA-Z0-9._-]{1,100}$`)

func validateRepoName(name string) error {
	if !validRepoName.MatchString(name) {
		return fmt.Errorf("repo name %q contains unsafe characters — check git remote origin", name)
	}
	return nil
}

// Keep this grammar aligned with config.validName: these values are forwarded
// to a remote Clem instance before its configuration is available locally.
var validAgentKey = regexp.MustCompile(`^[a-z][a-z0-9-]{0,30}$`)

func validateAgentKey(key string) error {
	if !validAgentKey.MatchString(key) {
		return fmt.Errorf("agent key must match %s, got %q", validAgentKey.String(), key)
	}
	return nil
}

// ValidatedRepoName returns RepoName after rejecting shell-metacharacter payloads.
func ValidatedRepoName() (string, error) {
	name, err := RepoName()
	if err != nil {
		return "", err
	}
	if err := validateRepoName(name); err != nil {
		return "", err
	}
	return name, nil
}
