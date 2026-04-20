# clem sample — Anthropic

Single-agent setup using Anthropic's Claude API directly. No local GPU or model
download needed — just an API key.

## Prerequisites

- An [Anthropic API key](https://console.anthropic.com/)
- A Discord bot token and server (see [Discord setup](../../docs/discord.md) if present, or the main README)
- A GitHub personal access token (for the agent to push code)

## Quick start (prebuilt image)

```bash
podman run -d --name clem-sample --systemd=always \
  -e DISCORD_SERVER_ID=... \
  -e DISCORD_GENERAL_CHANNEL=... \
  -e DISCORD_TASKS_CHANNEL=... \
  -e DISCORD_ALERTS_CHANNEL=... \
  -e DISCORD_LESSONS_CHANNEL=... \
  -p 7681:7681 \
  ghcr.io/jahwag/clem-sample:latest
```

Then provision the agent:

```bash
podman exec -it clem-sample bash
clem vault init
clem vault set anthropic ANTHROPIC_API_KEY=sk-ant-...
clem vault set github    GH_TOKEN=ghp_...
clem vault set discord-lead DISCORD_TOKEN=...
sudo clem provision
```

## Build from source

From the repo root:

```bash
docker build -f samples/Dockerfile --build-arg SAMPLE=anthropic -t clem-anthropic .
```

## Runtime modes

- **Tour** — interactive shell; explore `clem init` / `clem vault` without real credentials.
- **Runtime** — `--systemd=always` (podman) or `--privileged` (docker); `clem provision`
  creates OS users and starts services.

## Model

Defaults to `claude-sonnet-4-6`. Change the `model:` field in `clem.yaml` or override at
build time with `--build-arg SAMPLE=anthropic` and a customised `samples/anthropic/clem.yaml`.
