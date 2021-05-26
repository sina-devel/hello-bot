package actions

import (
	"fmt"
	"strings"

	gt "github.com/sina-devel/hello-bot/translategooglefree"
	tb "gopkg.in/tucnak/telebot.v2"
)

func (a *Actions) FaTranslatorHandler(m *tb.Message) {
	input := cmdRx.FindStringSubmatch(m.Text)[5]
	switch {
	case input == "" && m.IsReply():
		input = m.ReplyTo.Text
	case input == "":
		return
	}
	result, err := gt.Translate(input, "auto", "fa")
	if err != nil {
		return
	}
	text := strings.Builder{}
	for _, sentence := range result.Sentences {
		text.WriteString(sentence.Trans)
	}

	a.bot.Reply(
		m,
		fmt.Sprintf("from %s to fa\n%s", result.Src, text.String()),
	)
}
