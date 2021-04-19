package main

import (
	"github.com/sina-devel/hello-bot/actions"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
)

var (
	token     = os.Getenv("TOKEN")
	publicURL = os.Getenv("PUBLIC_URL")
	port      = os.Getenv("PORT")
)

func main() {
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

	botActions := actions.New(b)

	b.Handle(tb.OnAddedToGroup, botActions.OnUserJoinedHandler)
	b.Handle(tb.OnUserJoined, botActions.OnUserJoinedHandler)
	b.Handle(tb.OnDice, botActions.OnDiceHandler)
	b.Handle(tb.OnUserLeft, botActions.OnUserLeftHandler)
	b.Handle("/invite_link", botActions.InviteLinkHandler)
	b.Handle("/fa", botActions.FaTranslatorHandler)
	b.Handle("/en", botActions.EnTranslatorHandler)
	b.Handle("/dice", botActions.DiceHandler)
	b.Handle("pin it", botActions.PinHandler)
	b.Handle("unpin", botActions.UnpinHandler)

	b.Start()
}
