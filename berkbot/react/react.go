package berkbot_react

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Reaction func(session *discordgo.Session, message *discordgo.MessageCreate) (*discordgo.Emoji, *string)

var reactions = []Reaction{
	ReeReaction,
	KyleReaction,
	BrainReaction,
	MikeReaction,
	RipReaction,
}

func React(session *discordgo.Session, message *discordgo.MessageCreate) {
	for _, reaction := range reactions {
		emoji, str := reaction(session, message)
		var err error
		if str != nil {
			err = session.MessageReactionAdd(message.ChannelID, message.Message.ID, *str)
		}

		if emoji != nil {
			err = session.MessageReactionAdd(message.ChannelID, message.Message.ID, emoji.APIName())
		}

		if err != nil {
			log.Println("could not add emoji:", err)
		}
	}
}
