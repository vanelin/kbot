# kbot

DevOps application from scratch.
Simple CLI to send on-demand messages on behalf of a Telegram bot.

## Config / flags

URL to Telegram bot: [t.me/ivanvoloboyev_bot](https://t.me/ivanvoloboyev_bot) 

tested Go version = 1.19.7 

## Basic Installation
Steps on local env:
```bash
# password-protected SSH key
git clone git@github.com:vanelin/kbot.git 

cd kbot

# download and install dependency
go get

# type or copy-past telegram token for t.me/ivanvoloboyev_bot
read -s TELE_TOKEN

export TELE_TOKEN

# build kbot app
go build -ldflags "-X="github.com/vanelin/kbot/cmd.appVersion=v1.0.2

# start app
./kbot start
```
## Usage
After `./kbot start` go to Telegram bot and type **Srart**

Available commands:

1. `/start hello`
2. `/start ping`

