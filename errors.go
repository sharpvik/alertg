package alertg

import (
	"strings"

	"github.com/bigmate/slice"
	"gopkg.in/telebot.v4"
)

type ErrBroadcastFailed struct {
	ChatIDs []telebot.ChatID
}

func (ebf ErrBroadcastFailed) Error() string {
	return "failed to send message to: " + ebf.ids()
}

func (ebf ErrBroadcastFailed) AsError() error {
	if len(ebf.ChatIDs) > 0 {
		return ebf
	}

	return nil
}

func (ebf *ErrBroadcastFailed) Add(chatID telebot.ChatID) {
	ebf.ChatIDs = append(ebf.ChatIDs, chatID)
}

func (ebf ErrBroadcastFailed) ids() string {
	recipient := func(id telebot.ChatID) string { return id.Recipient() }
	recipients := slice.Map(ebf.ChatIDs, recipient)
	return strings.Join(recipients, ", ")
}
