package berkbot

import (
	berkbot_ping "github.com/bsssshhhhhhh/berkbot/berkbot/ping"
	berkbot_puns "github.com/bsssshhhhhhh/berkbot/berkbot/puns"
	berkbot_react "github.com/bsssshhhhhhh/berkbot/berkbot/react"
	"github.com/bwmarrin/discordgo"
)

func AddCommands(session *discordgo.Session) {
	session.AddHandler(berkbot_ping.Ping)
	session.AddHandler(berkbot_puns.Pun)
	session.AddHandler(berkbot_react.React)
}
