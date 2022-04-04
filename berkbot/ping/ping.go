package berkbot_ping

import (
	"github.com/bwmarrin/discordgo"
)

func Ping(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if message.Content == "!ping" {
		session.ChannelMessageSend(message.ChannelID, "pong")
	}
}
