package actions

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func (a *Actions) PinHandler(m *tb.Message) {
	if m.IsReply() {
		if err := a.bot.Pin(m.ReplyTo); err != nil {
			a.bot.Reply(m, "I can't ☹️")
		}
	} else {
		a.bot.Reply(m, "Are you ok? 🤔️")
	}
}
