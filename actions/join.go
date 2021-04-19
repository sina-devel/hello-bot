package actions

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

func (a *Actions) OnUserJoinedHandler(m *tb.Message) {
	a.bot.Reply(m, fmt.Sprintf("Hello %s %s 🖐️", m.UserJoined.FirstName, m.UserJoined.LastName))
}
