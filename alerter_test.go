package alertg_test

import (
	"testing"

	"github.com/sharpvik/alertg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/telebot.v4"
)

func TestInfo(t *testing.T) {
	var sender MockSender
	chatID := telebot.ChatID(int64(42))
	alt := alertg.Use(&sender).Notify(chatID)
	require.NoError(t, alt.Info("all systems operational"))
	assert.Equal(t, 1, sender.count)
}

type MockSender struct {
	count int
}

func (s *MockSender) Send(
	to telebot.Recipient,
	what any,
	opts ...any,
) (*telebot.Message, error) {
	s.count++
	return nil, nil
}
