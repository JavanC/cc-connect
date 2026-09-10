<p align="center">
  <img src="./docs/images/banner.svg" alt="CC-Connect Banner" width="800"/>
</p>

<p align="center">
  <a href="https://github.com/chenhg5/cc-connect/actions/workflows/ci.yml">
    <img src="https://github.com/chenhg5/cc-connect/actions/workflows/ci.yml/badge.svg" alt="CI Status"/>
  </a>
  <a href="https://github.com/chenhg5/cc-connect/releases">
    <img src="https://img.shields.io/github/v/release/chenhg5/cc-connect?include_prereleases" alt="Release"/>
  </a>
  <a href="https://www.npmjs.com/package/cc-connect">
    <img src="https://img.shields.io/npm/dm/cc-connect?logo=npm" alt="npm downloads"/>
  </a>
  <a href="https://github.com/chenhg5/cc-connect/blob/main/LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"/>
  </a>
  <a href="https://goreportcard.com/report/github.com/chenhg5/cc-connect">
    <img src="https://goreportcard.com/badge/github.com/chenhg5/cc-connect" alt="Go Report Card"/>
  </a>
</p>

<p align="center">
  <a href="https://discord.gg/kHpwgaM4kq">
    <img src="https://img.shields.io/badge/Discord-Join-5865F2?logo=discord&logoColor=white" alt="Discord"/>
  </a>
  <a href="https://t.me/+odGNDhCjbjdmMmZl">
    <img src="https://img.shields.io/badge/Telegram-Group-26A5E4?logo=telegram&logoColor=white" alt="Telegram"/>
  </a>
</p>

<p align="center">
  <a href="./README.md">English</a> | <a href="./README.zh-CN.md">中文</a>
</p>

<p align="center">
  <a href="https://trendshift.io/repositories/23266" target="_blank">
    <img src="https://trendshift.io/api/badge/repositories/23266" alt="chenhg5/cc-connect | Trendshift" style="width: 250px; height: 55px;" width="250" height="55"/>
  </a>
</p>



---

<br>

<p align="center">
  <b>Control your local AI agents from any chat app. Anywhere, anytime.</b>
</p>

<p align="center">
  cc-connect bridges AI agents running on your machine to the messaging platforms you already use.<br/>
  Code review, research, automation, data analysis — anything an AI agent can do,<br/>
  now accessible from your phone, tablet, or any device with a chat app.
</p>

<p align="center">
  <img src="docs/images/connector.png" alt="CC-Connect Architecture" width="90%"/>
</p>


## 🆕 What’s New in v1.5.1-beta.1

Beta since v1.5.0 stable — 16 merged PRs. Highlights:

- **i18n** — Localize agent system prompts (cron/timer/send/relay) based on language config (#1721).
- **Cursor** — Image attachments delivered via on-disk paths to the Cursor CLI (#1709).
- **Feishu** — Large file download via HTTP Range chunks, bypassing code=234037 (#1746); fail-closed when bot open_id discovery fails (#1725).
- **Weixin** — Reply and push paths now have separate send budgets (#1743); inbound dedup is configurable (#1733).
- **Claude Code** — `/compact` and slash commands restored by dropping `--replay-user-messages` (#1737); bounded session teardown (#1714).
- **Codex** — Failed app-server turns propagate (#1730); max reasoning effort supported (#1727); `/list` reads session names correctly (#1639).
- **Pi** — Attachments passed as `@path` refs (#1724); Windows build fix (#1738).

No breaking changes. See `changelogs/v1.5.1-beta.1.md` for the full changelog.


## 🧩 Platform feature snapshot

High-level view of what each **built-in platform** can do in cc-connect.

**Legend**

| Symbol | Meaning |
|--------|---------|
| ✅ | Works in **stable** cc-connect with typical configuration |
| ⚠️ | Partial, needs extra config (e.g. speech / ASR), or limited by the vendor app or API |
| ❌ | Not supported or not applicable in practice |

† **QQ (NapCat / OneBot)** — unofficial self-hosted bridge; behaviour depends on your NapCat / network setup.

| Capability | Feishu | WPS Xiezuo | DingTalk | Telegram | Slack | Discord | LINE | WeCom | Weibo | **Weixin**<br>*(personal)* | QQ† | QQ Bot | Matrix |
|------------|:------:|:----------:|:--------:|:--------:|:-----:|:-------:|:----:|:-----:|:-----:|:-------------------------:|:---:|:------:|:------:|
| Text & slash commands | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Markdown / cards | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ⚠️ | ⚠️ | ❌ | ✅ | ✅ | ✅ | ⚠️ |
| Streaming / chunked replies | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Images & files | ✅ | ❌ | ✅ | ✅ | ✅ | ✅ | ⚠️ | ✅ | ❌ | ✅ | ✅ | ✅ | ✅ |
| Voice / STT / TTS | ⚠️ | ❌ | ⚠️ | ✅ | ⚠️ | ⚠️ | ❌ | ⚠️ | ❌ | ✅ | ⚠️ | ⚠️ | ❌ |
| Private (DM) | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Group / channel | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ⚠️ | ✅ | ❌ | ✅ | ✅ | ✅ | ✅ |

> **WeCom:** Webhook mode needs a **public URL**; long-connection / WS style setups often do not.  
> **Voice row:** many platforms need `[speech]` / TTS providers enabled in `config.toml`; values are a best-effort summary.  
> Per-platform setup: [Platform setup guides](#-platform-setup-guides) below.


## ✨ Why cc-connect?

### 🤖 Universal Agent Support
**10+ AI Agents** — Claude Code, Codex, Cursor Agent, Kimi CLI, Qoder CLI, Gemini CLI, OpenCode, iFlow CLI, Pi, Devin, Copilot — plus any agent that supports the [Agent Client Protocol (ACP)](https://agentclientprotocol.com/get-started/agents). Use whichever fits your workflow, or all of them at once.

### 📱 Platform Flexibility
**13 Chat Platforms** — Feishu, WPS Xiezuo, DingTalk, Slack, Telegram, Discord, WeChat Work, Weibo, LINE, QQ, QQ Bot (Official), Matrix, plus **Weixin (personal ilink)** for **personal WeChat**. Most platforms need **zero public IP**.

### 🔄 Multi-Agent Orchestration
**Multi-Bot Relay** — Bind multiple bots in a group chat and let them communicate with each other. Ask Claude, get insights from Gemini — all in one conversation.

### 🎮 Complete Chat Control
**Full Control from Chat** — Switch models (`/model`), tune reasoning (`/reasoning`), change permission modes (`/mode`), manage sessions, all via slash commands.

**Directory Switching in Chat** — Change where the next session starts with `/dir <path>` (and `/cd <path>` as a compatibility alias), plus quick history jump via `/dir <number>` / `/dir -`.

### 🧠 Persistent Memory
**Agent Memory** — Read and write agent instruction files (`/memory`) without touching the terminal.

### ⏰ Intelligent Scheduling
**Scheduled Tasks** — Set up cron jobs in natural language. *"Every day at 6am, summarize GitHub trending"* just works.

### 🎤 Multimodal Support
**Voice & Images** — Send voice messages or screenshots; cc-connect handles STT/TTS and multimodal forwarding.

### 📦 Multi-Project Architecture
**Multi-Project** — One process, multiple projects, each with its own agent + platform combo.

### 🌍 Multilingual Interface
**5 Languages** — Native support for English, Chinese (Simplified & Traditional), Japanese, and Spanish. Built-in i18n ensures everyone feels at home.


<p align="center">
  <img src="docs/images/screenshot/cc-connect-lark.JPG" alt="飞书" width="32%" />
  <img src="docs/images/screenshot/cc-connect-telegram.JPG" alt="Telegram" width="32%" />
  <img src="docs/images/screenshot/cc-connect-wechat.JPG" alt="微信" width="32%" />
</p>
<p align="center">
  <em>Left：Lark &nbsp;|&nbsp; Telegram &nbsp;|&nbsp; Right：Wechat</em>
</p>


## 📋 Prerequisites

> **Install in this exact order** — cc-connect is a bridge for local AI coding agents, so the agent CLI must be installed and authenticated *before* cc-connect starts. Skipping ahead will cause `cc-connect` to exit with `claudecode: claude CLI not found in PATH` (or similar for your chosen agent), and the Web UI on `:9820` will never come up.

### 1️⃣ Install your AI Agent CLI

Pick the agent you want to bridge. You need **at least one**.

```bash
# Claude Code (most common)
brew install --cask claude-code            # macOS / Linux Homebrew
# or
npm install -g @anthropic-ai/claude-code   # any platform via npm

# OpenAI Codex
npm install -g @openai/codex

# Google Gemini CLI
npm install -g @google/gemini-cli

# iFlow CLI
npm install -g @iflow-ai/iflow-cli

# Qoder CLI
curl -fsSL https://qoder.com/install | bash
```

For **Cursor Agent** and **OpenCode**, follow the official install pages:
- Cursor Agent: <https://docs.cursor.com/agent>
- OpenCode: <https://github.com/opencode-ai/opencode>

Verify the binary is on your `PATH`:

```bash
claude --version       # or: codex / gemini / opencode / qodercli / cursor-agent ...
```

### 2️⃣ Authenticate the agent

Each agent has its own login flow — run the agent once interactively so it stores credentials in your home directory:

```bash
claude login           # opens a browser to authenticate
# or
codex login            # /gemini / opencode auth — see the agent's docs
```

If you skip this step, `cc-connect` will still start, but the agent will reject every prompt with an auth error.

### 3️⃣ Install cc-connect

```bash
# npm (any platform)
npm install -g cc-connect

# Homebrew (macOS / Linux)
brew install cc-connect

# Or download a binary from https://github.com/chenhg5/cc-connect/releases
```

### 4️⃣ Start cc-connect and open the Web UI

```bash
cc-connect             # starts the service; first run auto-creates ~/.cc-connect/config.toml
```

On first launch, cc-connect prints something like:

```
Web admin:  http://localhost:9820
```

Open that URL in your browser. If `9820` is already in use, pass `--web-port 9821` or set `web_port` in `config.toml`.

> **Note:** `cc-connect web` *only* opens the browser and the config UI — it does **not** start the service. You still need `cc-connect` running in another terminal.

### 5️⃣ Configure platform bot tokens in the Web UI

In the Web UI, create a project, then add at least one platform (Feishu / Telegram / Discord / Slack / DingTalk / WeChat Work / QQ / LINE / Weixin) and paste the bot token from that platform's developer console. Save and cc-connect will hot-reload.

That's it — send a message to your bot and cc-connect will relay it to your local agent.

---

## 🚀 Quick Start

### 🤖 Install & Configure via AI Agent (Recommended)

> **The easiest way** — Send this to Claude Code or any AI coding agent, and it will handle the entire installation and configuration for you:

```bash
Follow https://raw.githubusercontent.com/chenhg5/cc-connect/refs/heads/main/INSTALL.md to install and configure cc-connect.
```


### 📦 Manual Install

**Via npm:**

```bash
npm install -g cc-connect
```

**Via Homebrew (macOS / Linux):**

```bash
brew install cc-connect
```

**Download binary from [GitHub Releases](https://github.com/chenhg5/cc-connect/releases):**

```bash
# Linux amd64 - Stable
curl -L -o cc-connect https://github.com/chenhg5/cc-connect/releases/latest/download/cc-connect-linux-amd64
chmod +x cc-connect
sudo mv cc-connect /usr/local/bin/

```

**Build from source (requires Go 1.22+):**

```bash
git clone https://github.com/chenhg5/cc-connect.git
cd cc-connect
make build
```


### ⚙️ Configure

> **💡 Tip: Use the Web UI to configure** — After installing, run `cc-connect web` to configure the web admin and open the dashboard in your browser. You can visually create projects, add platforms, manage providers, and chat with your agent — no need to manually edit TOML files. **Note:** `cc-connect web` only configures and opens the browser — you still need to run `cc-connect` separately to start the service.

If you prefer manual configuration:

```bash
mkdir -p ~/.cc-connect
cp config.example.toml ~/.cc-connect/config.toml
vim ~/.cc-connect/config.toml
```

Set `admin_from = "alice,bob"` in a project to allow those user IDs to run privileged commands such as `/dir` and `/shell`.
`admin_from` must be placed under `[[projects]]` (not under `[projects.platforms.options]`). You can use `/whoami` or `/status` to get your current `User ID`.
When a user runs `/dir reset`, cc-connect restores the configured `work_dir` and clears the persisted override stored under `data_dir/projects/<project>.state.json`.


### ▶️ Run

```bash
./cc-connect
```


### 🔄 Upgrade

```bash
# npm
npm install -g cc-connect

# Homebrew
brew upgrade cc-connect

# Binary self-update
cc-connect update           # Stable
cc-connect update --pre     # Include pre-releases
```


## 📊 Support Matrix

| Component | Type | Status |
|-----------|------|--------|
| Agent | Claude Code | ✅ Supported |
| Agent | Codex (OpenAI) | ✅ Supported |
| Agent | Cursor Agent | ✅ Supported |
| Agent | Gemini CLI (Google) | ✅ Supported |
| Agent | Qoder CLI | ✅ Supported |
| Agent | OpenCode (Crush) | ✅ Supported |
| Agent | iFlow CLI | ✅ Supported |
| Agent | Kimi CLI (Moonshot) | ✅ Supported |
| Agent | Pi (Cursor Background Agent) | ✅ Supported |
| Agent | Copilot (GitHub) | ✅ Supported |
| Agent | ACP (Agent Client Protocol) | ✅ Any [ACP-compatible agent](https://agentclientprotocol.com/get-started/agents) |
| Agent | Devin (Cognition) | ✅ Supported (via ACP) |
| Agent | Goose (Block) | 🔜 Planned |
| Agent | Aider | 🔜 Planned |
| Platform | Feishu (Lark) | ✅ WebSocket — no public IP needed |
| Platform | DingTalk | ✅ Stream — no public IP needed |
| Platform | WPS Xiezuo | ✅ WebSocket — no public IP needed |
| Platform | Telegram | ✅ Long Polling — no public IP needed |
| Platform | Slack | ✅ Socket Mode — no public IP needed |
| Platform | Discord | ✅ Gateway — no public IP needed |
| Platform | Weibo | ✅ WebSocket — no public IP needed |
| Platform | LINE | ✅ Webhook — public URL required |
| Platform | WeChat Work | ✅ WebSocket / Webhook |
| Platform | Weixin (personal, ilink) | ✅— HTTP long polling — no public IP needed |
| Platform | QQ (NapCat/OneBot) | ✅ WebSocket |
| Platform | QQ Bot (Official) | ✅ WebSocket — no public IP needed |
| Platform | Matrix | ✅ Long Polling (/sync) — no public IP needed |


## 📖 Platform Setup Guides

| Platform | Guide | Connection | Public IP? |
|----------|-------|------------|------------|
| Feishu (Lark) | [docs/feishu.md](docs/feishu.md) | WebSocket | No |
| DingTalk | [docs/dingtalk.md](docs/dingtalk.md) | Stream | No |
| WPS Xiezuo | [docs/wps-xiezuo.md](docs/wps-xiezuo.md) | WebSocket | No |
| Telegram | [docs/telegram.md](docs/telegram.md) | Long Polling | No |
| Slack | [docs/slack.md](docs/slack.md) | Socket Mode | No |
| Google Chat | [docs/googlechat.md](docs/googlechat.md) | Cloud Pub/Sub | No |
| Discord | [docs/discord.md](docs/discord.md) | Gateway | No |
| Weibo | [docs/weibo.md](docs/weibo.md) | WebSocket | No |
| WeChat Work | [docs/wecom.md](docs/wecom.md) | WebSocket / Webhook | No (WS) / Yes (Webhook) |
| Weixin (personal) | [docs/weixin.md](docs/weixin.md) | HTTP long polling (ilink) | No |
| QQ / QQ Bot | [docs/qq.md](docs/qq.md) | WebSocket | No |
| Matrix | [docs/matrix.md](docs/matrix.md) | /sync (Long Polling) | No |


## 🎯 Key Features

### 💬 Session Management

```
/new [name]       Start a new session
/list             List all sessions
/switch <id>      Switch session
/current          Show current session
/dir [path|reset] Show, switch, or reset work directory
```

Project configs rotate to a fresh session automatically after long inactivity. This prevents "context drift" where stale chat history (failed commands, debugging noise) is repeatedly re-ingested via `--continue` and starts to dominate the model's attention. The previous session is preserved and remains accessible via `/list` and `/switch`.

```toml
[[projects]]
reset_on_idle_mins = 30   # default when unset; set to 0 to disable
```

The default is **30 minutes** when unset. Set `reset_on_idle_mins = 0` to opt out and always continue the previous session.

### 🛡️ OS-User Isolation (`run_as_user`)

On Linux/macOS, a project can spawn its agent under a different Unix
user for OS-level file-system isolation from the supervisor user that
runs cc-connect. Currently supported by Claude Code.

```toml
[[projects]]
name = "claude-sandboxed"
run_as_user = "partseeker-coder"
run_as_env = ["PGSSLROOTCERT"]
```

The target user needs passwordless sudo from the supervisor, no sudo
of its own, read+write on `work_dir`, and its own `~/.claude/settings.json`
with whatever credentials the agent uses. If you authenticate via
`claude.ai` OAuth, symlink the target user's `~/.claude/.credentials.json`
to the supervisor's copy so token refresh stays in sync — see the
[environment propagation checklist](./docs/usage.md#environment-propagation-what-moves-into-the-target-users-home)
for details. See
[`docs/usage.md`](./docs/usage.md#running-agents-as-a-different-unix-user-run_as_user)
for the full setup.

Before starting cc-connect, audit the setup with:

```bash
cc-connect doctor user-isolation
```

This runs three go/no-go preflight gates and an isolation probe that
reports what the target user can and cannot read. cc-connect refuses to
start if any gate fails or if the probe detects a cross-user leak.

---

### 🔐 Permission Modes

```
/mode             Show available modes
/mode yolo        # Auto-approve all tools
/mode default     # Ask for each tool
```


### 🔄 Provider Management

```
/provider list              List providers
/provider switch <name>     Switch API provider at runtime
```


### 🤖 Model Selection

```
/model                      List available models (format: alias - model)
/model switch <alias>       Switch to model by alias
```


### 📂 Work Directory

```
/dir                         Show current work directory and history
/dir <path>                  Switch to a path (relative or absolute)
/dir <number>                Switch from history
/dir -                       Switch to previous directory
/cd <path>                   Compatibility alias for /dir <path>
```


### ⏰ Scheduled Tasks

```bash
/cron add 0 6 * * * Summarize GitHub trending
```

### 📎 Agent Attachment Send-Back

When an agent generates a local screenshot, chart, PDF, bundle, or other file, it can send that attachment back to the current chat.

First release supports:
- Feishu
- Telegram

If your agent does not natively inject the system prompt, run this once in chat after upgrading:

```text
/bind setup
```

or:

```text
/cron setup
```

This refreshes the cc-connect instructions in the project memory file so the agent knows how to send attachments back.

You can control this feature globally in `config.toml`:

```toml
attachment_send = "on"  # default: "on"; set to "off" to block image/file send-back
```

This switch is independent from the agent's `/mode`. It only controls `cc-connect send --image/--file`. Voice send-back uses the TTS config instead.

Examples:

```bash
cc-connect send --image /absolute/path/to/chart.png
cc-connect send --file /absolute/path/to/report.pdf
cc-connect send --file /absolute/path/to/report.pdf --image /absolute/path/to/chart.png
cc-connect send --tts "Hello from cc-connect"
```

Notes:
- Absolute paths are the safest option.
- `--image` and `--file` can both be repeated.
- `--tts` sends synthesized speech when the user asks for a voice reply.
- `attachment_send = "off"` disables only attachment send-back; ordinary text replies still work.
- Attachments are capped at 50 MiB by default; configure with `max_attachment_size_mb` (or `CC_MAX_ATTACHMENT_SIZE_MB` env, same MiB unit).
- This command is for generated attachments, not ordinary text replies.

📖 **Full documentation:** [docs/usage.md](docs/usage.md)


## 📚 Documentation

- [Usage Guide](docs/usage.md) — Complete feature documentation
- [INSTALL.md](INSTALL.md) — AI-agent-friendly installation guide
- [config.example.toml](config.example.toml) — Configuration template
- [CONTRIBUTING.md](CONTRIBUTING.md) — How to report issues and contribute pull requests


## 👥 Community

- [Discord](https://discord.gg/kHpwgaM4kq)
- [Telegram](https://t.me/+odGNDhCjbjdmMmZl)


## ☕ Support the Project

If cc-connect has been helpful to you, consider buying us a coffee! Your support helps us:

- 🛠️ Maintain and improve the project
- 📚 Write better documentation and tutorials
- 🐛 Fix bugs and add new features faster
- ☕ Keep the developers caffeinated

### How to Donate

**Buy Me a Coffee**: [https://buymeacoffee.com/cg33](https://buymeacoffee.com/cg33)

**WeChat Pay / Alipay**:

| WeChat Pay | Alipay |
|:----------:|:------:|
| <img src="docs/images/wechatpay.jpg" alt="WeChat Pay" width="150"> | <img src="docs/images/alipay.jpg" alt="Alipay" width="150"> |

### Thank You, Donors! 🎉

We're grateful to everyone who has supported this project. Leave your GitHub username in the donation message if you'd like to be recognized here!

<!-- Donors will be listed below -->
| Avatar | GitHub Username | Date |
|--------|-----------------|------|
| <img src="https://avatars.githubusercontent.com/u/1762560?v=4" width="40" height="40" style="border-radius: 50%;"> | [@thx0701](https://github.com/thx0701) | 2026-04-29 |


## 🤝 Commercial Cooperation

We accept the following commercial collaborations:

- **Enterprise Customization**: Custom deployment for internal AI tooling (Feishu, DingTalk, WeChat Work, Slack, etc.)
- **Technical Consulting**: AI agent integration and architecture design
- **Outsourcing Projects**: AI-related system development

**Contact**: **Email**: chg80333@gmail.com | **WeChat**: mongorz | [Telegram](https://t.me/+odGNDhCjbjdmMmZl) | [Discord](https://discord.gg/kHpwgaM4kq)


## 🙏 Contributors

<a href="https://github.com/chenhg5/cc-connect/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=chenhg5/cc-connect&v=20250313" />
</a>


## ⭐ Star History

<a href="https://www.star-history.com/#chenhg5/cc-connect&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=chenhg5/cc-connect&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=chenhg5/cc-connect&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=chenhg5/cc-connect&type=Date" />
 </picture>
</a>


## 📄 License

MIT License


<p align="center">
  <sub>Built with ❤️ by the cc-connect community</sub>
</p>
