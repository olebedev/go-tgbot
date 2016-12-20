package tgbot

import (
	"fmt"

	"github.com/olebedev/go-tgbot/models"
	"github.com/pkg/errors"
)

type session string

func (s session) String() string {
	return string(s)
}

// GetSession is a default session repository. it's just
// extract a text from an update. Returns Stringer interface
// to be able to match the update.
func GetSession(u *models.Update) (fmt.Stringer, error) {
	var text string
	switch true {
	case u.Message != nil:
		text = u.Message.Text
	case u.EditedMessage != nil:
		text = u.EditedMessage.Text
	case u.ChannelPost != nil:
		text = u.ChannelPost.Text
	case u.EditedChannelPost != nil:
		text = u.EditedChannelPost.Text
	case u.InlineQuery != nil:
		text = u.InlineQuery.Query
	case u.CallbackQuery != nil:
		text = u.CallbackQuery.Data
	case u.ChosenInlineResult != nil:
		if u.ChosenInlineResult.ResultID != nil {
			text = *u.ChosenInlineResult.ResultID
		}
	default:
		return nil, errors.New("unknown update")
	}

	return session(text), nil
}
