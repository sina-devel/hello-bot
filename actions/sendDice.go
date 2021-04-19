package actions

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"math/rand"
	"time"
)

func (a *Actions) DiceHandler(m *tb.Message) {
	dices := []*tb.Dice{tb.Ball, tb.Goal, tb.Slot, tb.Dart, tb.Cube, {Type: "🎳"}}
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	a.bot.Reply(m, dices[rnd.Intn(len(dices))])
}
