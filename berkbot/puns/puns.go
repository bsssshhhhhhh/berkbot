package berkbot_puns

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

var puns = []string{
	"tbd",
}

func GetRandomPun() string {
	randomIndex := rand.Intn(len(puns))
	return puns[randomIndex]
}

func Pun(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!pun" {
		s.ChannelMessageSend(m.ChannelID, GetRandomPun())
	}
}
