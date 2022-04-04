package berkbot_react

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var BrainEmoji string = "🧠"

func BrainReaction(session *discordgo.Session, message *discordgo.MessageCreate) (*discordgo.Emoji, *string) {
	match, _ := regexp.MatchString("brain", message.Content)

	if match {
		return nil, &BrainEmoji
	}

	return nil, nil
}
