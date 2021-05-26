package actions

import (
	"math/rand"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func (a *Actions) OnDiceHandler(m *tb.Message) {
	successes := []string{"😎", "😍", "🤠", "🤩", "🙂"}
	failures := []string{"🤕", "🙁", "😶", "😑", "😭"}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	success := successes[random.Intn(len(failures))]
	failure := failures[random.Intn(len(failures))]

	time.Sleep(4 * time.Second)

	switch m.Dice.Type {
	case tb.Cube.Type:
		switch m.Dice.Value {
		case 6:
			a.bot.Reply(m, success)
		default:
			a.bot.Reply(m, failure)
		}
	case tb.Ball.Type:
		switch m.Dice.Value {
		case 4, 5:
			a.bot.Reply(m, success)
		default:
			a.bot.Reply(m, failure)
		}
	case tb.Goal.Type:
		switch m.Dice.Value {
		case 3, 4, 5:
			a.bot.Reply(m, success)
		default:
			a.bot.Reply(m, failure)
		}
	case tb.Dart.Type:
		switch m.Dice.Value {
		case 6:
			a.bot.Reply(m, success)
		default:
			a.bot.Reply(m, failure)
		}
	case tb.Slot.Type:
		switch m.Dice.Value {
		case 64:
			a.bot.Reply(m, success)
		default:
			a.bot.Reply(m, failure)
		}
	case "🎳":
		switch m.Dice.Value {
		case 6:
			a.bot.Reply(m, success)
		default:
			a.bot.Reply(m, failure)
		}
	}
}
