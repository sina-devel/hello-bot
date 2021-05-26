package actions

import (
	"regexp"

	tb "gopkg.in/tucnak/telebot.v2"
)

var cmdRx = regexp.MustCompile(`^(/\w+)(@(\w+))?(\s|$)(?s)(.+)?`)

type Actions struct {
	bot *tb.Bot
}

func New(bot *tb.Bot) *Actions {
	return &Actions{
		bot: bot,
	}
}
