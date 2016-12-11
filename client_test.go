package tgbot

import (
	"context"
	"os"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/go-tgbot/client/messages"
	"github.com/stretchr/testify/require"
)

var token = ""
var uid = ""

func init() {
	token = os.Getenv("TOKEN")
	uid = os.Getenv("USER_ID")
}

func TestGetMe(t *testing.T) {
	cli := NewClient(context.Background(), token)
	r, err := cli.Users.GetMe(nil)
	require.Nil(t, err)
	require.Equal(t, true, *r.Payload.Ok)
}

func TestSendMessage(t *testing.T) {
	cli := NewClient(context.Background(), token)
	r, err := cli.Messages.SendMessage(
		messages.NewSendMessageParams().
			WithRequest(messages.SendMessageBody{
				ChatID: uid,
				Text:   pointer.ToString("test message"),
			}),
	)
	require.Nil(t, err)
	require.Equal(t, true, *r.Payload.Ok)
	require.Equal(t, "test message", r.Payload.Result.Text)
}
