package actions

import (
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

			u.Rights = tb.NoRights()

			if err := a.bot.Restrict(m.Chat, u); err != nil {
				a.bot.Reply(m, "I can't change permission")
			}
		}
	}
}
