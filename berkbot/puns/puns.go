package berkbot_puns

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

var puns = []string{
	"tbd",
}

func getRandomPun() string {
	return puns[rand.Intn(len(puns))]
}

func Pun(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!pun" {
		s.ChannelMessageSend(m.ChannelID, getRandomPun())
	}
}
