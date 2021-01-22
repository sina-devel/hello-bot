package main

import (
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")       // sets automatically
		publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
		token     = os.Getenv("TOKEN")      // you must add it to your config vars
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.SetCommands([]tb.Command{
		{
			Text:        "invitelink",
			Description: "send group invitelink",
		},
		{
			Text:        "dice",
			Description: "roll the dice",
		},
	})

	b.Handle(tb.OnAddedToGroup, func(m *tb.Message) {
		b.Reply(m, "🙂️ سلام خوش اومدی عزیز")
	})

	b.Handle("/invitelink", func(m *tb.Message) {
		if inviteLink, err := b.GetInviteLink(m.Chat); err == nil {
			linkmsg, _ := b.Reply(m, inviteLink)
			go func(m *tb.Message, lm *tb.Message) {
				<-time.NewTimer(5 * time.Minute).C
				b.Delete(m)
				b.Delete(lm)
			}(m, linkmsg)
		} else {
			if m.Chat.Type != tb.ChatGroup || m.Chat.Type == tb.ChatSuperGroup {
				b.Reply(m, "اینجا گروه خصوصی نی 😅️🤣️")
			} else {
				b.Reply(m, "منم مثل تو نمیدونم 😅️🤣️")
			}
		}
	})

	b.Handle("/dice", func(m *tb.Message) {
		b.Reply(m, "\xF0\x9F\x8E\xB2")
	})

	b.Handle(tb.OnUserLeft, func(m *tb.Message) {
		b.Reply(m, "بری دیگه برنگردی😏️")
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		if m.Text == "اینو پین کن" {
			if m.IsReply() {
				if err := b.Pin(m.ReplyTo); err != nil {
					b.Reply(m, "نمیتونم پین کنم ☹️")
				}
			} else {
				b.Reply(m, "چی رو پین کنم دقیقا 🤔️")
			}
		}
	})

	b.Start()
}
