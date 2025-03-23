# alertg

Go library to issue alerts through [Telegram Messenger](https://telegram.org/).

## Why

In my personal projects, I found it convenient to receive system alerts through
Telegram. I don't have to check monitoring dashboards. Instead, updates are
delivered to me directly.

## Get

```sh
go get github.com/sharpvik/alertg
```

## Setup

### Create the bot

We assume that [you have created a Telegram bot](https://core.telegram.org/bots/tutorial).
You will need its token to initialise `telebot` (see next section).

### Find chat ID

To find your chat ID, send any message to your bot, then open this link in your
browser:

```text
https://api.telegram.org/bot<TOKEN>/getUpdates
                            ^^^^^^^
         REPLACE THIS WITH YOUR SPECIFIC TELEGRAM BOT TOKEN
```

The chat ID can be found in the resulting JSON.

> We support broadcasting alerts to multiple chats. If you want to do that, you
> should use the aforementioned method to get chat ID of every person wanting
> to receive system alerts.

## Use

```go
token := "YOUR TELEGRAM BOT TOKEN"
chatID := telebot.ChatID(int64(42)) // YOUR CHAT ID
sender, _ := telebot.NewBot(telebot.Settings{Token: token})
altg := alertg.Use(sender).Notify(chatID) // ONE OR MORE CHAT ID
altg.Info("all systems operational")
```

> Example above omits error handling for brevity but you shouldn't.

## Interface

We support the following alert levels (inspired by [log/slog](https://pkg.go.dev/log/slog)).

```go
altg.Debug("message")
altg.Info("message")
altg.Warn("message")
altg.Error("message")
```
