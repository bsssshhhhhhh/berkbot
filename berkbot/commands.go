package berkbot

import (
	berkbot_ping "github.com/bsssshhhhhhh/berkbot/berkbot/ping"
	"github.com/bwmarrin/discordgo"
)

func AddCommands(s *discordgo.Session) {
	s.AddHandler(berkbot_ping.Ping)
}
