package berkbot_react

import (
	"github.com/bwmarrin/discordgo"
)

type Reaction func(s *discordgo.Session, m *discordgo.MessageCreate) *discordgo.Emoji

var reactions = []Reaction{ReeReaction}

func React(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, reaction := range reactions {
		emoji := reaction(s, m)
		if emoji != nil {
			s.MessageReactionAdd(m.ChannelID, m.Message.ID, emoji.ID)
		}
	}
}
