package berkbot_react

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var BrainEmoji string = "🧠"

func BrainReaction(s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.Emoji, *string) {
	match, _ := regexp.MatchString("brain", m.Content)

	if match {
		return nil, &BrainEmoji
	}

	return nil, nil
}
