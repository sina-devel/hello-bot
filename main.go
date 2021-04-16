package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	gt "github.com/sina-devel/hello-bot/translategooglefree"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		token     = os.Getenv("TOKEN")
		publicURL = os.Getenv("PUBLIC_URL")
		port      = os.Getenv("PORT")
		cmdRx     = regexp.MustCompile(`^(/\w+)(@(\w+))?(\s|$)(?s)(.+)?`)
	)

	pref := tb.Settings{
		Token: token,
		Poller: &tb.Webhook{
			Listen: ":" + port,
			Endpoint: &tb.WebhookEndpoint{
				PublicURL: publicURL,
			},
		},
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	if err := b.SetCommands([]tb.Command{
		{
			Text:        "invite_link",
			Description: "send group inviteLink",
		},
		{
			Text:        "dice",
			Description: "roll the dice",
		},
		{
			Text:        "fa",
			Description: "translation text to persian",
		},
		{
			Text:        "en",
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
		success := "oh my god, you are very lucky"
		failure := "I have not seen anyone more unlucky than you 🤣"
		time.Sleep(2 * time.Second)
		switch m.Dice.Type {
		case tb.Cube.Type:
			switch m.Dice.Value {
			case 6:
				_, _ = b.Reply(m, success)
			default:
				_, _ = b.Reply(m, failure)
			}
		case tb.Ball.Type, tb.Goal.Type:
			switch m.Dice.Value {
			case 4, 5:
				_, _ = b.Reply(m, success)
			default:
				_, _ = b.Reply(m, failure)
			}
		case tb.Dart.Type:
			switch m.Dice.Value {
			case 6:
				_, _ = b.Reply(m, success)
			default:
				_, _ = b.Reply(m, failure)
			}
		case tb.Slot.Type:
			switch m.Dice.Value {
			case 64:
				_, _ = b.Reply(m, success)
			default:
				_, _ = b.Reply(m, failure)
			}
		case "🎳":
			switch m.Dice.Value {
			case 6:
				_, _ = b.Reply(m, success)
			default:
				_, _ = b.Reply(m, failure)
			}
		}
		if _, err := b.Reply(m, fmt.Sprintf("value of Dice: %d", m.Dice.Value)); err != nil {
			log.Println(err)
		}
	})

	b.Handle("/invite_link", func(m *tb.Message) {
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

	b.Handle("/fa", func(m *tb.Message) {
		input := cmdRx.FindStringSubmatch(m.Text)[5]
		if m.IsReply() && input == "" {
			input = m.ReplyTo.Text
		}
		result, err := gt.Translate(input, "auto", "fa")
		if err != nil {
			if _, err := b.Reply(m, err.Error()); err != nil {
				log.Println(err)
			}
			return
		}
		text := strings.Builder{}
		for _, sentence := range result.Sentences {
			text.WriteString(sentence.Trans)
		}

		if _, err := b.Reply(m, result, fmt.Sprintf("from %s to fa\n%s", result.Src, text.String())); err != nil {
			log.Println(err)
		}
	})

	b.Handle("/en", func(m *tb.Message) {
		input := cmdRx.FindStringSubmatch(m.Text)[5]
		if m.IsReply() && input == "" {
			input = m.ReplyTo.Text
		}
		result, err := gt.Translate(input, "auto", "en")
		if err != nil {
			if _, err := b.Reply(m, err.Error()); err != nil {
				log.Println(err)
			}
			return
		}
		text := strings.Builder{}
		for _, sentence := range result.Sentences {
			text.WriteString(sentence.Trans)
		}

		if _, err := b.Reply(m, fmt.Sprintf("from %s to en\n%s", result.Src, text.String())); err != nil {
			log.Println(err)
		}
	})

	b.Handle("/dice", func(m *tb.Message) {
		dices := []*tb.Dice{tb.Ball, tb.Goal, tb.Slot, tb.Dart, tb.Cube, {Type: "🎳"}}
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

	b.Handle("pin it", func(m *tb.Message) {
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
	})

	b.Handle("unpin", func(m *tb.Message) {
		if err := b.Unpin(m.Chat); err != nil {
			if _, err := b.Reply(m, "I can't ☹️"); err != nil {
				log.Println(err)
			}
		}
	})

	b.Start()
}
