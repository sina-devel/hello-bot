package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"

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

	if err := b.SetCommands([]tb.Command{
		{
			Text:        "inviteLink",
			Description: "send group inviteLink",
		},
		{
			Text:        "dice",
			Description: "roll the dice",
		},
		{
			Text:        "toFA",
			Description: "translation text to persian",
		},
		{
			Text:        "toEN",
			Description: "translation text to english",
		},
	}); err != nil {
		log.Fatalln(err)
	}

	b.Handle(tb.OnAddedToGroup, func(m *tb.Message) {
		if _, err := b.Reply(m, fmt.Sprintf("Hello %s %s 🖐️", m.UserJoined.FirstName, m.UserJoined.LastName)); err != nil {
			log.Println(err)
		}
	})

	b.Handle(tb.OnUserJoined, func(m *tb.Message) {
		if _, err := b.Reply(m, fmt.Sprintf("Hello %s %s 🖐️", m.UserJoined.FirstName, m.UserJoined.LastName)); err != nil {
			log.Println(err)
		}
	})

	b.Handle(tb.OnDice, func(m *tb.Message) {
		if _, err := b.Reply(m, fmt.Sprintf("value of Dice: %d", m.Dice.Value)); err != nil {
			log.Println(err)
		}
	})

	b.Handle("/inviteLink", func(m *tb.Message) {
		if inviteLink, err := b.GetInviteLink(m.Chat); err == nil {
			lm, _ := b.Reply(m, inviteLink)
			go func(m *tb.Message, lm *tb.Message) {
				<-time.NewTimer(5 * time.Minute).C
				_ = b.Delete(m)
				_ = b.Delete(lm)
			}(m, lm)
		} else {
			if m.Chat.Type != tb.ChatGroup {
				if _, err := b.Reply(m, "link 404 😅️🤣️"); err != nil {
					log.Println(err)
				}
			} else {
				if _, err := b.Reply(m, "I don't know like you 😅️🤣️"); err != nil {
					log.Println(err)
				}
			}
		}
	})

	b.Handle("/toFA", func(m *tb.Message) {
		pat := regexp.MustCompile(`^/toFA(@[a-zA-Z0-9_]*)?[\s]*(.*)`)
		text := pat.ReplaceAllString(m.Text, "$2")
		result, err := gt.Translate(text, "auto", "fa")
		if err != nil {
			if _, err := b.Reply(m, err.Error()); err != nil {
				log.Println(err)
			}
			return
		}
		if _, err := b.Reply(m, result); err != nil {
			log.Println(err)
		}
	})

	b.Handle("/toEN", func(m *tb.Message) {
		pat := regexp.MustCompile(`^/toEN(@[a-zA-Z0-9_]*)?[\s]*(.*)`)
		text := pat.ReplaceAllString(m.Text, "$2")
		result, err := gt.Translate(text, "auto", "en")
		if err != nil {
			if _, err := b.Reply(m, err.Error()); err != nil {
				log.Println(err)
			}
			return
		}
		if _, err := b.Reply(m, result); err != nil {
			log.Println(err)
		}
	})
	b.Handle("/dice", func(m *tb.Message) {
		dices := []*tb.Dice{tb.Cube, tb.Dart, tb.Ball, tb.Goal, tb.Slot}
		rnd := rand.New(rand.NewSource(time.Now().Unix()))
		if _, err := b.Reply(m, dices[rnd.Intn(len(dices))]); err != nil {
			log.Println(err)
		}
	})

	b.Handle(tb.OnUserLeft, func(m *tb.Message) {
		if _, err := b.Reply(m, fmt.Sprintf("GoodBye %s", m.UserLeft.FirstName)); err != nil {
			log.Println(err)
		}
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		if m.Text == "pin it" {
			if m.IsReply() {
				if err := b.Pin(m.ReplyTo); err != nil {
					if _, err := b.Reply(m, "I can't ☹️"); err != nil {
						log.Println(err)
					}
				}
			} else {
				if _, err := b.Reply(m, "Are you ok? 🤔️"); err != nil {
					log.Println(err)
				}
			}
		}
		if m.Text == "unpin" {
			if err := b.Unpin(m.Chat); err != nil {
				if _, err := b.Reply(m, "I can't ☹️"); err != nil {
					log.Println(err)
				}
			}
		}
	})

	b.Start()
}
