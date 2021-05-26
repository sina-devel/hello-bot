package actions

import (
	"fmt"

	tb "gopkg.in/tucnak/telebot.v2"
)

func (a *Actions) OnUserLeftHandler(m *tb.Message) {
	a.bot.Reply(m, fmt.Sprintf("GoodBye %s", m.UserLeft.FirstName))
}
