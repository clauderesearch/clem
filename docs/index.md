---
layout: default
title: clem — Continuously Looping Engineering Machines
description: docker-compose for Claude Code agents. Persistent teams of agents, 24/7, on any Linux host.
---

<p align="center">
  <img src="logo.png" alt="clem" width="180">
</p>

<h1 align="center" style="margin-bottom:0">clem</h1>

<p align="center"><em>Continuously Looping Engineering Machines.</em></p>

<p align="center">
  <a href="https://github.com/jahwag/clem/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-MIT-green?style=flat-square" alt="MIT License"></a>
  <a href="https://github.com/jahwag/clem/releases"><img src="https://img.shields.io/github/v/release/jahwag/clem?style=flat-square&color=orange" alt="Latest release"></a>
  <a href="https://goreportcard.com/report/github.com/jahwag/clem"><img src="https://goreportcard.com/badge/github.com/jahwag/clem?style=flat-square" alt="Go Report Card"></a>
  <a href="https://discord.gg/pR4qeMH4u4"><img src="https://img.shields.io/badge/Discord-community-5865F2?style=flat-square&logo=discord&logoColor=white" alt="Discord"></a>
</p>

---

`clem` runs a team of Claude Code agents 24/7 on any Linux host. Each agent is a separate OS user in a tmux session under systemd. Agents coordinate over Discord or Slack, pick up tasks, write code, and open PRs. A watchdog restarts anything that crashes. Configure once, walk away.

## Install

```bash
curl -fsSL https://github.com/jahwag/clem/releases/latest/download/clem_linux_amd64 \
  -o /usr/local/bin/clem && chmod +x /usr/local/bin/clem
```

Other platforms + checksums on the [releases page](https://github.com/jahwag/clem/releases/latest).

## Quickstart

```bash
mkdir my-agents && cd my-agents
clem init                  # writes clem.yaml + CLAUDE.local.md template
clem vault init            # generates age keypair, creates .sops.yaml
clem vault set discord DISCORD_TOKEN=...
sudo clem provision        # creates OS users, tmux, systemd, watchdog
sudo clem login lead       # one-time OAuth per agent
sudo clem up               # agents live 24/7
```

Full README with config reference, coordination backends (Discord / Slack), runtime backends (claude-code / opencode), provider options (Anthropic / Bedrock / Vertex / Ollama / OpenAI-compat) lives on [GitHub](https://github.com/jahwag/clem#readme).

## Features

| | |
|---|---|
| **Per-agent OS identity** | Each agent = its own Linux user, own git identity, own bot account. Crash boundaries are real. |
| **Multi-coordination** | Discord + Slack today via one config knob. |
| **Multi-runtime** | `claude-code` or `opencode`. Mix Anthropic cloud, Bedrock, Vertex, Ollama, OpenAI-compat. |
| **Encrypted secrets** | Per-agent `.env` materialised from age/sops vaults at provision time. |
| **Self-healing** | systemd + tmux + watchdog. Stalled sessions get restarted. |
| **Bring your own model** | Default Claude; one flag away from Ollama / Bedrock / Vertex / local. Tested on NVIDIA Nemotron. |
| **Live ops** | `clem status` for health; ttyd web terminal per agent if you want to watch. |
| **No Kubernetes** | Runs on a laptop, home server, Raspberry Pi, or small VPS. |

## Why

Because an agent team shouldn't require a cluster. Because running Claude Code sessions by hand doesn't scale past a day. Because Linux already has everything you need — OS users, tmux, systemd — and the right abstraction is just *configure it once*.

## Community

- [GitHub Discussions](https://github.com/jahwag/clem/discussions)
- [Discord server](https://discord.gg/pR4qeMH4u4)
- [Issues](https://github.com/jahwag/clem/issues) — bug reports + feature requests

Licensed under the [MIT License](https://github.com/jahwag/clem/blob/main/LICENSE).
