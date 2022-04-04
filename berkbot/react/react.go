package berkbot_react

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Reaction func(s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.Emoji, *string)

var reactions = []Reaction{
	ReeReaction,
	KyleReaction,
	BrainReaction,
}

func React(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, reaction := range reactions {
		emoji, str := reaction(s, m)
		var err error
		if str != nil {
			err = s.MessageReactionAdd(m.ChannelID, m.Message.ID, *str)
		}

		if emoji != nil {
			err = s.MessageReactionAdd(m.ChannelID, m.Message.ID, emoji.APIName())
		}

		if err != nil {
			fmt.Println("could not add emoji:", err)
		}
	}
}
