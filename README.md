# Kbot

A simple Telegram bot built in Go using the Cobra CLI framework and Telebot v4 for handling Telegram messages. 
---

## Prerequisites
- Go 1.16 or later
- Telegram Bot Token (set as TELE_TOKEN environment variable)
- Required Go packages:
    - github.com/spf13/cobra
    - gopkg.in/telebot.v4

## Installation

```bash
# Clone the repository
git clone https://github.com/denis-ruban/kbot.git
cd kbot

# Build the binary
go build -ldflags "-X github.com/denis-ruban/kbot/cmd.AppVersion=v1.0.0" -o kbot

# Run the bot via environment variable
export TELE_TOKEN="<YOUR_TELEGRAM_BOT_TOKEN>"
```
---
## Usage
Start the bot:
```bash
./kbot start
```

**Commands:**
- `/start` – welcome message  
- `/hello`, `/hi` – friendly greeting  
- `/help` – list available commands  
- `/version` – show current bot version  
- `/echo <text>` – repeat back any text  