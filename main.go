package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
        "strings"

	gt "github.com/sina-devel/hello-bot/translategooglefree"
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
		{
			Text:        "tofa",
			Description: "translation text to persian",
		},
	})

	b.Handle(tb.OnAddedToGroup, func(m *tb.Message) {
		b.Reply(m, fmt.Sprintf("Hello %s %s 🖐️", m.UserJoined.FirstName, m.UserJoined.LastName))
	})

	b.Handle(tb.OnUserJoined, func(m *tb.Message) {
		b.Reply(m, fmt.Sprintf("Hello %s %s 🖐️", m.UserJoined.FirstName, m.UserJoined.LastName))
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
				b.Reply(m, "link 404 😅️🤣️")
			} else {
				b.Reply(m, "I don't know like you 😅️🤣️")
			}
		}
	})

	b.Handle("/tofa", func(m *tb.Message) {
                text := strings.ReplaceAll(m.Text, "/tofa", "")
		result, err := gt.Translate(text, "auto", "fa")
		if err != nil {
			b.Reply(m, "error 😰")
			return
		}
		b.Reply(m, result)
	})

	b.Handle("/dice", func(m *tb.Message) {
		dices := []*tb.Dice{tb.Cube, tb.Dart, tb.Ball, tb.Goal, tb.Slot}
		rnd := rand.New(rand.NewSource(time.Now().Unix()))
		b.Reply(m, dices[rnd.Intn(len(dices))])
	})

	b.Handle(tb.OnUserLeft, func(m *tb.Message) {
		b.Reply(m, fmt.Sprintf("GoodBye %s", m.UserLeft.FirstName))
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		if m.Text == "pin it" {
			if m.IsReply() {
				if err := b.Pin(m.ReplyTo); err != nil {
					b.Reply(m, "I can't ☹️")
				}
			} else {
				b.Reply(m, "Are you ok? 🤔️")
			}
		}
		if m.Text == "unpin" {
			if err := b.Unpin(m.Chat); err != nil {
				b.Reply(m, "I can't ☹️")
			}
		}
	})

	b.Start()
}
