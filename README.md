# BMTool 
![](https://img.shields.io/badge/license-MIT-blue)
![](https://img.shields.io/badge/GO-1.17-blue)
![](https://img.shields.io/badge/NodeJS-v16-green)
![](https://img.shields.io/badge/PRs-welcome-green)

> This project does not have auth. If you need to deploy publicly, please ensure security (add auth for secondary development)

Simple expenses or income, statistics percentage, view monthly bill records. Support Telegram Bot, cli and WebAPI.

- `/` (Root directory): Server
- `tool-cli`: Working on the command line
- `telegram-bot`: Working on the telegram bot

# Start using

If you don't want to build it yourself, you can download zip/tar.gz in the [Release](https://github.com/ArsFy/bill_management_tool/releases)

## Server

### Build and Run

```bash
# Clone Repo
git clone https://github.com/ArsFy/bill_management_tool.git
cd bill_management_tool

# Build
go mod tidy
go build .

# Run
chmod 777 bmtool_server
./bmtool_server
```

#### Config

Rename `config.example.json` to `config.json`

```json
{ 
    "port": "80"
}
```

## Cli Tool

### Build and Run

```bash
cd bill_management_tool/tool-cli

# Build
go mod tidy
go build .

# Run
chmod 777 tool-cli
./tool-cli
```

## Telegram Bot

### Config

Rename `config.example.json` to `config.json`

```js
{
    "server": "https://example.com",      // BMTool Server
    "bot_token": "123456:ABCDEFGHIGKLMN", // Telegram Bot Token @BotFather
    "admin": [ 123456789 ]                // Admin's (yours) User ID (Number)
}
```

### Run

```bash
cd bill_management_tool/telegram-bot

# Run
npm i
node main.js
```

# What's this?
This project helps you to better manage your team's money, count and make summaries. It is an easy productivity tool that can be used by teams.