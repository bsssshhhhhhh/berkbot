package berkbot

import (
	berkbot_ping "github.com/bsssshhhhhhh/berkbot/berkbot/ping"
	berkbot_puns "github.com/bsssshhhhhhh/berkbot/berkbot/puns"
	"github.com/bwmarrin/discordgo"
)

func AddCommands(s *discordgo.Session) {
	s.AddHandler(berkbot_ping.Ping)
	s.AddHandler(berkbot_puns.Pun)
}
