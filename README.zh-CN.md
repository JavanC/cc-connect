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
  <b>在任何聊天工具里，远程操控你的本地 AI Agent。随时随地，随心所欲。</b>
</p>

<p align="center">
  cc-connect 把运行在你机器上的 AI Agent 桥接到你日常使用的即时通讯工具。<br/>
  代码审查、资料研究、自动化任务、数据分析 —— 只要 AI Agent 能做的事，<br/>
  都能通过手机、平板或任何有聊天应用的设备来完成。
</p>

<p align="center">
  <img src="docs/images/connector.png" alt="CC-Connect 架构图" width="90%"/>
</p>


## 🆕 v1.5.1-beta.1 更新了什么

自 v1.5.0 正式版以来的 beta —— 16 个已合并 PR。亮点：

- **i18n** — 按 language 配置本地化 agent system prompt（cron/timer/send/relay）(#1721)。
- **Cursor** — 图片附件通过 on-disk path 传给 Cursor CLI (#1709)。
- **飞书** — 大文件 HTTP Range 分块下载，绕过 code=234037 (#1746)；bot open_id 发现失败 fail-closed (#1725)。
- **微信** — 回复/推送分路径 send budget (#1743)；入站 dedup 可配置 (#1733)。
- **Claude Code** — 移除 `--replay-user-messages`，恢复 `/compact` 等 slash 命令 (#1737)；session teardown 竞态修复 (#1714)。
- **Codex** — 传播 app-server turn 失败 (#1730)；支持 max reasoning effort (#1727)；`/list` 正确读 session 名 (#1639)。
- **Pi** — 附件 @path 引用而非 inline bytes (#1724)；Windows 编译修复 (#1738)。

无任何破坏性变更。完整 changelog 见 `changelogs/v1.5.1-beta.1.md`。


## 🧩 平台能力一览

内置各渠道在 cc-connect 里的大致能力对照，方便快速对比。

**图例**

| 符号 | 含义 |
|------|------|
| ✅ | **稳定版** cc-connect + 常规配置下可用 |
| ⚠️ | 部分支持、需额外配置（如语音/STT）或受厂商接口 / 应用类型限制 |
| ❌ | 不支持或实际不可用 |

† **QQ（NapCat / OneBot）** — 非官方自建桥接，体验依赖你的 NapCat 与网络环境。

| 能力 | 飞书 | WPS 协作 | 钉钉 | Telegram | Slack | Discord | LINE | 企业微信 | 微博 | **微信个人号**<br>（ilink） | QQ† | QQ 官方机器人 | Matrix |
|------|:----:|:--------:|:----:|:--------:|:-----:|:-------:|:----:|:--------:|:----:|:--------------------------:|:---:|:------------:|:-----:|
| 文本与斜杠命令 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Markdown / 卡片 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ⚠️ | ⚠️ | ❌ | ✅ | ✅ | ✅ | ⚠️ |
| 流式 / 分片回复 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| 图片与文件 | ✅ | ❌ | ✅ | ✅ | ✅ | ✅ | ⚠️ | ✅ | ❌ | ✅ | ✅ | ✅ | ✅ |
| 语音 / STT / TTS | ⚠️ | ❌ | ⚠️ | ✅ | ⚠️ | ⚠️ | ❌ | ⚠️ | ❌ | ✅ | ⚠️ | ⚠️ | ❌ |
| 私聊 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| 群聊 / 频道 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ⚠️ | ✅ | ❌ | ✅ | ✅ | ✅ | ✅ |

> **企业微信：** Webhook 模式需要**公网 URL**；长连接等模式多数**不需要**。  
> **语音行：** 多数平台要在 `config.toml` 里配置 `[speech]` / TTS 等，表中为经验性归纳。  
> 分平台接入步骤见下文 [平台接入指南](#-平台接入指南)。


## ✨ 为什么选择 cc-connect？

### 🤖 通用 Agent 支持
**10+ 大 AI Agent** — Claude Code、Codex、Cursor Agent、Kimi CLI、Qoder CLI、Gemini CLI、OpenCode、iFlow CLI、Pi、Devin、Copilot，还可通过 [Agent Client Protocol (ACP)](https://agentclientprotocol.com/get-started/agents) 接入更多 Agent。按需选用，或同时使用全部。

### 📱 平台灵活性
**13 大聊天平台** — 飞书、WPS 协作、钉钉、Slack、Telegram、Discord、企业微信、微博、LINE、QQ、QQ 官方机器人、Matrix，以及 **微信个人号（ilink）**。大部分平台**无需公网 IP**。

### 🔄 多 Agent 编排
**多机器人中继** — 在群聊中绑定多个机器人，让它们相互协作。问 Claude，再听 Gemini 的见解 — 同一个对话搞定。

### 🎮 完整的聊天控制
**聊天即控制** — 切换模型 (`/model`)、切换推理强度 (`/reasoning`)、切换权限模式 (`/mode`)、管理会话，全部通过斜杠命令完成。

**聊天切换工作目录** — 使用 `/dir <路径>` 切换下一次会话启动目录（`/cd <路径>` 为兼容别名），并支持 `/dir <序号>` / `/dir -` 快速在历史目录间跳转。

### 🧠 持久化记忆
**Agent 记忆** — 在聊天中直接读写 Agent 指令文件 (`/memory`)，无需回到终端。

### ⏰ 智能定时任务
**定时任务** — 自然语言创建 cron 任务。"每天早上6点总结 GitHub trending" 即刻生效。

### 🎤 多模态支持
**语音 & 图片** — 发语音或截图，cc-connect 自动处理 STT/TTS 和多模态转发。

### 📦 多项目架构
**多项目管理** — 一个进程同时管理多个项目，各自独立的 Agent + 平台组合。

### 🌍 多语言界面
**5 种语言** — 原生支持英语、中文（简体/繁体）、日语和西班牙语。内置 i18n 让每个人都能得心应手。


<p align="center">
  <img src="docs/images/screenshot/cc-connect-lark.JPG" alt="飞书" width="32%" />
  <img src="docs/images/screenshot/cc-connect-telegram.JPG" alt="Telegram" width="32%" />
  <img src="docs/images/screenshot/cc-connect-wechat.JPG" alt="微信" width="32%" />
</p>
<p align="center">
  <em>左：飞书 &nbsp;|&nbsp; Telegram &nbsp;|&nbsp; 右：微信</em>
</p>


## 📋 准备工作

> **请严格按照以下顺序安装** — cc-connect 是本地 AI 编程 Agent 的桥接工具，因此对应的 Agent CLI 必须先安装并完成登录认证，之后 cc-connect 才能正常启动。如果跳过前面的步骤直接启动 cc-connect，进程会直接退出并报错 `claudecode: claude CLI not found in PATH`（其他 Agent 报错类似），Web UI 在 `:9820` 也就无从访问。

### 1️⃣ 安装 AI Agent CLI

选择你要桥接的 Agent，至少装一个。

```bash
# Claude Code（最常用）
brew install --cask claude-code            # macOS / Linux Homebrew
# 或
npm install -g @anthropic-ai/claude-code   # 任意平台通过 npm

# OpenAI Codex
npm install -g @openai/codex

# Google Gemini CLI
npm install -g @google/gemini-cli

# iFlow CLI
npm install -g @iflow-ai/iflow-cli

# Qoder CLI
curl -fsSL https://qoder.com/install | bash
```

**Cursor Agent** 和 **OpenCode** 请参考各自的官方安装文档：
- Cursor Agent: <https://docs.cursor.com/agent>
- OpenCode: <https://github.com/opencode-ai/opencode>

确认可执行文件在 `PATH` 中：

```bash
claude --version       # 或 codex / gemini / opencode / qodercli / cursor-agent ...
```

### 2️⃣ 完成 Agent 登录认证

每个 Agent 都有自己的登录流程 — 先在终端交互式跑一次，让它把凭据存到你的 home 目录：

```bash
claude login           # 会在浏览器里打开授权页面
# 或
codex login            # / gemini / opencode 等也类似，请参考各自文档
```

跳过这一步的话，cc-connect 仍能启动，但 Agent 会因为未认证拒绝所有请求。

### 3️⃣ 安装 cc-connect

```bash
# npm（任意平台）
npm install -g cc-connect

# Homebrew（macOS / Linux）
brew install cc-connect

# 也可以从 https://github.com/chenhg5/cc-connect/releases 直接下载二进制
```

### 4️⃣ 启动 cc-connect 并打开 Web UI

```bash
cc-connect             # 启动服务；首次运行会自动生成 ~/.cc-connect/config.toml
```

首次启动时，cc-connect 会打印类似：

```
Web admin:  http://localhost:9820
```

在浏览器里打开该地址。如果 `9820` 已被占用，可以传 `--web-port 9821` 或在 `config.toml` 里设置 `web_port`。

> **注意：** `cc-connect web` *只* 打开浏览器和配置界面，并**不会**启动服务本身。仍需要在另一个终端里跑 `cc-connect`。

### 5️⃣ 在 Web UI 里配置平台 Bot Token

在 Web UI 里新建一个项目，然后添加至少一个平台（飞书 / Telegram / Discord / Slack / 钉钉 / 企业微信 / QQ / LINE / 微信 ilink），把该平台开发者后台的 Bot Token 粘贴进去。保存后 cc-connect 会热加载。

至此完成 — 给你的 Bot 发条消息，cc-connect 就会把它转给本地的 Agent。

---

## 🚀 快速开始

### 🤖 通过 AI Agent 安装配置（推荐）

> **最简单的方式** — 把这段话发给 Claude Code 或其他 AI 编码 Agent，它会帮你完成整个安装和配置过程：

```bash
请参考 https://raw.githubusercontent.com/chenhg5/cc-connect/refs/heads/main/INSTALL.md 帮我安装和配置 cc-connect
```


### 📦 手动安装

**通过 npm：**

```bash
# npm install -g cc-connect
```

**通过 Homebrew（macOS / Linux）：**

```bash
brew install cc-connect
```

**从 [GitHub Releases](https://github.com/chenhg5/cc-connect/releases) 下载：**

```bash
# Linux amd64 - 稳定版
curl -L -o cc-connect https://github.com/chenhg5/cc-connect/releases/latest/download/cc-connect-linux-amd64
chmod +x cc-connect
sudo mv cc-connect /usr/local/bin/

```

**从源码编译（需要 Go 1.22+）：**

```bash
git clone https://github.com/chenhg5/cc-connect.git
cd cc-connect
make build
```


### ⚙️ 配置

> **💡 推荐使用 Web UI 配置** — 安装完成后，运行 `cc-connect web` 配置 Web 管理后台并在浏览器中打开。可以可视化创建项目、添加平台、管理服务商、直接和 Agent 聊天，无需手动编辑 TOML 文件。**注意：** `cc-connect web` 仅用于配置和打开浏览器，并不会启动 cc-connect 服务本身，你仍需单独运行 `cc-connect` 来启动。

如果你更喜欢手动配置：

```bash
mkdir -p ~/.cc-connect
cp config.example.toml ~/.cc-connect/config.toml
vim ~/.cc-connect/config.toml
```

在项目配置里设置 `admin_from = "alice,bob"` 后，只有这些用户 ID 才能执行 `/dir`、`/shell` 等特权命令。
执行 `/dir reset` 时，cc-connect 会恢复配置中的 `work_dir`，并清除保存在 `data_dir/projects/<project>.state.json` 里的目录覆盖状态。


### ▶️ 运行

```bash
./cc-connect
```


### 🔄 升级

```bash
# npm
npm install -g cc-connect

# Homebrew
brew upgrade cc-connect

# 二进制自更新
cc-connect update           # 稳定版
cc-connect update --pre     # 含预发布版本
```


## 📊 支持状态

| 组件 | 类型 | 状态 |
|------|------|------|
| Agent | Claude Code | ✅ 已支持 |
| Agent | Codex (OpenAI) | ✅ 已支持 |
| Agent | Cursor Agent | ✅ 已支持 |
| Agent | Gemini CLI (Google) | ✅ 已支持 |
| Agent | Qoder CLI | ✅ 已支持 |
| Agent | OpenCode (Crush) | ✅ 已支持 |
| Agent | iFlow CLI | ✅ 已支持 |
| Agent | Kimi CLI (Moonshot) | ✅ 已支持 |
| Agent | Pi (Cursor Background Agent) | ✅ 已支持 |
| Agent | Copilot (GitHub) | ✅ 已支持 |
| Agent | ACP (Agent Client Protocol) | ✅ 支持任何 [ACP 兼容 Agent](https://agentclientprotocol.com/get-started/agents) |
| Agent | Devin (Cognition) | ✅ 已支持（通过 ACP）|
| Agent | Goose (Block) | 🔜 计划中 |
| Agent | Aider | 🔜 计划中 |
| Platform | 飞书 (Lark) | ✅ WebSocket — 无需公网 IP |
| Platform | 钉钉 | ✅ Stream — 无需公网 IP |
| Platform | WPS 协作 | ✅ WebSocket — 无需公网 IP |
| Platform | Telegram | ✅ Long Polling — 无需公网 IP |
| Platform | Slack | ✅ Socket Mode — 无需公网 IP |
| Platform | Discord | ✅ Gateway — 无需公网 IP |
| Platform | 微博 | ✅ WebSocket — 无需公网 IP |
| Platform | LINE | ✅ Webhook — 需要公网 URL |
| Platform | 企业微信 | ✅ WebSocket / Webhook |
| Platform | 微信个人号（ilink） | ✅— HTTP 长轮询 — 无需公网 IP |
| Platform | QQ (NapCat/OneBot) | ✅ WebSocket |
| Platform | QQ 官方机器人 | ✅ WebSocket — 无需公网 IP |
| Platform | Matrix | ✅ Long Polling (/sync) — 无需公网 IP |


## 📖 平台接入指南

| 平台 | 指南 | 连接方式 | 需要公网 IP? |
|------|------|---------|-------------|
| 飞书 (Lark) | [docs/feishu.md](docs/feishu.md) | WebSocket | 不需要 |
| 钉钉 | [docs/dingtalk.md](docs/dingtalk.md) | Stream | 不需要 |
| WPS 协作 | [docs/wps-xiezuo.md](docs/wps-xiezuo.md) | WebSocket | 不需要 |
| Telegram | [docs/telegram.md](docs/telegram.md) | Long Polling | 不需要 |
| Slack | [docs/slack.md](docs/slack.md) | Socket Mode | 不需要 |
| Discord | [docs/discord.md](docs/discord.md) | Gateway | 不需要 |
| 微博 | [docs/weibo.md](docs/weibo.md) | WebSocket | 不需要 |
| 企业微信 | [docs/wecom.md](docs/wecom.md) | WebSocket / Webhook | 不需要 (WS) / 需要 (Webhook) |
| 微信个人号（ilink） | [docs/weixin.md](docs/weixin.md) | HTTP 长轮询（ilink） | 不需要 |
| QQ / QQ 机器人 | [docs/qq.md](docs/qq.md) | WebSocket | 不需要 |
| Matrix | [docs/matrix.md](docs/matrix.md) | /sync（长轮询） | 不需要 |


## 🎯 核心功能

### 💬 会话管理

```
/new [名称]            创建新会话
/list                  列出所有会话
/switch <id>           切换会话
/current               查看当前会话
/dir [路径|reset]      查看、切换或重置工作目录
```

项目配置也可以开启“长时间空闲后自动切到新会话”：

```toml
[[projects]]
reset_on_idle_mins = 60
```


### 🛡️ 系统用户隔离 (`run_as_user`)

在 Linux/macOS 上，项目可以用另一个 Unix 用户身份启动 Agent，从而在操作系统层面实现文件系统隔离。目前 Claude Code 已支持。

```toml
[[projects]]
name = "claude-sandboxed"
run_as_user = "partseeker-coder"
run_as_env = ["PGSSLROOTCERT"]
```

目标用户需要：supervisor 对其配置免密 sudo、自身不拥有 sudo、对 `work_dir` 有读写权限、拥有自己的 `~/.claude/settings.json`。
如果你通过 `claude.ai` OAuth 认证，请将目标用户的 `~/.claude/.credentials.json` 软链接到 supervisor 的副本以保持 token 同步 —— 详见[环境传播清单](./docs/usage.md#environment-propagation-what-moves-into-the-target-users-home)。
完整设置说明见 [`docs/usage.md`](./docs/usage.md#running-agents-as-a-different-unix-user-run_as_user)。

启动 cc-connect 之前，可用以下命令审核配置：

```bash
cc-connect doctor user-isolation
```

该命令会执行三项前置检查和一次隔离探测，报告目标用户能/不能读取的内容。如果任一检查失败或探测到跨用户泄漏，cc-connect 将拒绝启动。

---

### 🔐 权限模式

```
/mode             查看可用模式
/mode yolo        # 自动批准所有工具
/mode default     # 每次工具调用前询问
```


### 🔄 Provider 管理

```
/provider list              列出 Provider
/provider switch <名称>     运行时切换 API Provider
```


### 🤖 模型选择

```
/model                      列出可用模型（格式：alias - model）
/model switch <alias>       按别名切换模型
```


### 📂 工作目录

```
/dir                         查看当前工作目录与历史
/dir <路径>                  切换到指定目录（相对或绝对路径）
/dir <序号>                  按历史序号切换
/dir -                       返回上一个目录
/cd <路径>                   `/dir <路径>` 的兼容别名
```


### ⏰ 定时任务

```bash
/cron add 0 6 * * * 帮我总结 GitHub trending
```

### 📎 Agent 回传图片和文件

当 Agent 在本地生成了截图、图表、PDF、日志包等文件时，可以主动把附件发回当前聊天。

首版支持：
- 飞书
- Telegram

如果当前 Agent 不是原生注入 system prompt 的类型，升级后请先在聊天里执行一次：

```text
/bind setup
```

或：

```text
/cron setup
```

这样会把最新的 cc-connect 指令写入项目记忆文件，Agent 才会知道如何回传附件。

你也可以在 `config.toml` 里全局控制这项能力：

```toml
attachment_send = "on"  # 默认 "on"；设为 "off" 会禁用图片/文件回传
```

这个开关与 agent 的 `/mode` 独立，只控制 `cc-connect send --image/--file` 这条附件回传路径。语音回传走 TTS 配置。

回传方式：

```bash
cc-connect send --image /absolute/path/to/chart.png
cc-connect send --file /absolute/path/to/report.pdf
cc-connect send --file /absolute/path/to/report.pdf --image /absolute/path/to/chart.png
cc-connect send --tts "你好"
```

要点：
- 使用绝对路径最稳妥。
- `--image` 和 `--file` 都可以重复传多个。
- `--tts` 用当前 TTS provider 合成并发送语音，适合用户自然要求“发语音”的场景。
- `attachment_send = "off"` 只会关闭附件回传，普通文本回复仍然正常。
- 单个附件默认上限 50 MiB；可用 `max_attachment_size_mb` 配置（或环境变量 `CC_MAX_ATTACHMENT_SIZE_MB`，同样单位 MiB）。
- 这个命令是给“生成后的附件回传”用的，不是给普通文本回复用的。

📖 **完整文档：** [docs/usage.zh-CN.md](docs/usage.zh-CN.md)


## 📚 文档

- [使用指南](docs/usage.zh-CN.md) — 完整功能文档
- [INSTALL.md](INSTALL.md) — AI Agent 友好的安装指南
- [config.example.toml](config.example.toml) — 配置模板
- [CONTRIBUTING.md](CONTRIBUTING.md) — Issue / PR 提交流程与贡献说明


## 👥 社区

- [Discord](https://discord.gg/kHpwgaM4kq)
- [Telegram](https://t.me/+odGNDhCjbjdmMmZl)


## ☕ 支持项目

如果 cc-connect 对你有帮助，请考虑请我们喝杯咖啡！你的支持帮助我们：

- 🛠️ 维护和改进项目
- 📚 编写更好的文档和教程
- 🐛 更快修复 bug 和添加新功能
- ☕ 让开发者保持精力充沛

### 捐赠方式

**Buy Me a Coffee**：[https://buymeacoffee.com/cg33](https://buymeacoffee.com/cg33)

**微信支付 / 支付宝**：

| 微信支付 | 支付宝 |
|:----------:|:------:|
| <img src="docs/images/wechatpay.jpg" alt="微信支付" width="150"> | <img src="docs/images/alipay.jpg" alt="支付宝" width="150"> |

### 感谢捐赠者！🎉

感谢每一位支持这个项目的朋友。捐赠时留言你的 GitHub 用户名，我们会在这里展示！

<!-- 捐赠者名单 -->
| 头像 | GitHub 用户名 | 日期 |
|------|-----------------|------|
| <img src="https://avatars.githubusercontent.com/u/1762560?v=4" width="40" height="40" style="border-radius: 50%;"> | [@thx0701](https://github.com/thx0701) | 2026-04-29 |


## 🤝 商业合作

我们接受以下商业合作：

- **企业定制**：为企业定制内部 AI 工具入口（飞书、钉钉、企业微信、Slack 等）
- **技术咨询**：AI agent 集成方案设计与架构咨询
- **外包项目**：AI 相关系统开发

**联系方式**：**邮箱**：chg80333@gmail.com | **微信**：mongorz | [Telegram](https://t.me/+odGNDhCjbjdmMmZl) | [Discord](https://discord.gg/kHpwgaM4kq)


## 🙏 贡献者

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
  <sub>由 cc-connect 社区用 ❤️ 构建</sub>
</p>
