package actions

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

func (a *Actions) InviteLinkHandler(m *tb.Message) {
	if inviteLink, err := a.bot.GetInviteLink(m.Chat); err == nil {
		lm, _ := a.bot.Reply(m, inviteLink)
		go func(m *tb.Message, lm *tb.Message) {
			<-time.NewTimer(5 * time.Minute).C
			a.bot.Delete(m)
			a.bot.Delete(lm)
		}(m, lm)
	} else {
		if m.Chat.Type != tb.ChatGroup {
			if _, err := a.bot.Reply(m, "link 404 😅️🤣️"); err != nil {
				log.Println(err)
			}
		} else {
			if _, err := a.bot.Reply(m, "I don't know like you 😅️️🤣"); err != nil {
				log.Println(err)
			}
		}
	}
}
