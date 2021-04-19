package actions

import (
	"fmt"
	gt "github.com/sina-devel/hello-bot/translategooglefree"
	tb "gopkg.in/tucnak/telebot.v2"
	"strings"
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
		a.bot.Reply(m, err.Error())
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