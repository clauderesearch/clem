#!/usr/bin/env bash
# clem end-to-end provision demo, designed for asciinema recording.
#
# Requirements: clem, age, sops, yq on PATH; sudo access.
# Works on a bare Linux host or inside samples/Dockerfile (--privileged, /bin/bash entrypoint).
#
# Record:
#   asciinema rec demo.cast -- bash docs/demo.sh
#   asciinema upload demo.cast          # copy the cast ID into README.md + docs/index.html
#
# Tokens are placeholders: provisioning completes cleanly; agents need real tokens to run.
# Regenerate after clem changes: re-run on a fresh host and re-upload.

set -euo pipefail

DEMO_DIR=$(mktemp -d /tmp/clem-demo-XXXX)
trap 'rm -rf "$DEMO_DIR"' EXIT

run() {
    printf '\033[1;32m$\033[0m %s\n' "$*"
    sleep 0.4
    eval "$*"
    sleep 1
}

clear
printf '# clem — docker-compose for Claude Code agents\n'
printf '# Provisioning a two-agent team from scratch\n\n'
sleep 2

cd "$DEMO_DIR"

run "clem init"
run "clem vault init"

# Placeholder tokens — provisioning flow completes; agents fail at runtime (expected for demo)
run "clem vault set github GH_TOKEN=ghp_fake"
run "clem vault set discord-lead DISCORD_TOKEN=discord_fake"

run "sudo clem provision"
run "clem status"

printf '\n# Agents provisioned. Start them: sudo clem up\n'
