package alertg

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

type Sender interface {
	Send(to telebot.Recipient, what any, opts ...any) (*telebot.Message, error)
}

type Alerter struct {
	sender  Sender
	chatIDs []telebot.ChatID
}

func Use(sender Sender) *Alerter {
	return &Alerter{
		sender: sender,
	}
}

func (alt *Alerter) Notify(chatIDs ...telebot.ChatID) *Alerter {
	alt.chatIDs = chatIDs
	return alt
}

func (alt *Alerter) Debug(format string, args ...any) error {
	msg := fmt.Sprintf("ℹ️ "+format, args...)
	return alt.broadcast(msg)
}

func (alt *Alerter) Info(format string, args ...any) error {
	msg := fmt.Sprintf("📰 "+format, args...)
	return alt.broadcast(msg)
}

func (alt *Alerter) Warn(format string, args ...any) error {
	msg := fmt.Sprintf("⚠️ "+format, args...)
	return alt.broadcast(msg)
}

func (alt *Alerter) Error(format string, args ...any) error {
	msg := fmt.Sprintf("🚨 "+format, args...)
	return alt.broadcast(msg)
}

func (alt *Alerter) broadcast(msg string) error {
	var failures ErrBroadcastFailed

	for _, chatID := range alt.chatIDs {
		if _, err := alt.sender.Send(chatID, msg); err != nil {
			failures.Add(chatID)
		}
	}

	return failures.AsError()
}
