package actions

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

func (a *Actions) OnDiceHandler(m *tb.Message) {
	success := "oh my god, you are very lucky"
	failure := "I have not seen anyone more unlucky than you 🤣"
	time.Sleep(2 * time.Second)
	switch m.Dice.Type {
	case tb.Cube.Type:
		switch m.Dice.Value {
		case 6:
			a.bot.Reply(m, success)
		default:
			a.bot.Reply(m, failure)
		}
	case tb.Ball.Type, tb.Goal.Type:
		switch m.Dice.Value {
		case 4, 5:
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
