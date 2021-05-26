package actions

import tb "gopkg.in/tucnak/telebot.v2"

func isAdmin(id int, admins []tb.ChatMember) bool {
	for _, admin := range admins {
		if id == admin.User.ID {
			return true
		}
	}
	return false
}
