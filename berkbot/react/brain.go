package berkbot_react

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func BrainReaction(s *discordgo.Session, m *discordgo.MessageCreate) *discordgo.Emoji {
	match, _ := regexp.MatchString("brain", m.Content)

	if match {
		guild, err := s.Guild(m.GuildID)
		if err != nil {
			fmt.Println("Could not get guild "+m.GuildID+": ", err)
			return nil
		}

		for _, emoji := range guild.Emojis {
			if emoji.Name == "brain" {
				return emoji
			}
		}
	}
	return nil
}
