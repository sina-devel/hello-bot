package actions

import (
	"strconv"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func (a *Actions) MuteHandler(m *tb.Message) {
	if m.IsReply() &&
		m.Chat.Type == tb.ChatGroup ||
		m.Chat.Type == tb.ChatSuperGroup {

		admins, err := a.bot.AdminsOf(m.Chat)
		if err != nil {
			a.bot.Reply(m, "I can't")
		}

		if isAdmin(m.Sender.ID, admins) {
			u, err := a.bot.ChatMemberOf(m.Chat, m.ReplyTo.Sender)
			if err != nil {
				a.bot.Reply(m, "WTF")
			}

			if until, err := strconv.ParseInt(m.Payload, 10, 64); err == nil {
				u.RestrictedUntil = int64(time.Duration(until) * time.Minute)
			}

			u.Rights = tb.NoRestrictions()

			if err := a.bot.Restrict(m.Chat, u); err != nil {
				a.bot.Reply(m, "I can't change permission")
			}
		}
	}
}
