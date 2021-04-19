package actions

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func (a *Actions) UnpinHandler(m *tb.Message) {
	if err := a.bot.Unpin(m.Chat); err != nil {
		a.bot.Reply(m, "I can't ☹️")
	}
}
