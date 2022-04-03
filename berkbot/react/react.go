package berkbot_react

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Reaction func(s *discordgo.Session, m *discordgo.MessageCreate) *discordgo.Emoji

var reactions = []Reaction{
	ReeReaction,
	KyleReaction,
	BrainReaction,
}

func React(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, reaction := range reactions {
		emoji := reaction(s, m)
		if emoji != nil {
			err := s.MessageReactionAdd(m.ChannelID, m.Message.ID, emoji.ID)

			if err != nil {
				fmt.Println("Could not react to message")
			}
		}
	}
}
