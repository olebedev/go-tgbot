package tgbot

import (
	"fmt"
	"net/url"

	"github.com/olebedev/go-tgbot/models"
	"github.com/pkg/errors"
)

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
		if u.ChosenInlineResult.Query != nil {
			text = *u.ChosenInlineResult.Query
		}
	default:
		return nil, errors.New("unknown update")
	}

	return url.Parse(text)
}
