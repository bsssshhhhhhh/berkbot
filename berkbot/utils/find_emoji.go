package berkbot_utils

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func FindGuildEmoji(session *discordgo.Session, guildID string, emojiName string) *discordgo.Emoji {
	guild, err := session.Guild(guildID)
	if err != nil {
		log.Println("Could not get guild " + guildID)
		return nil
	}

	for _, emoji := range guild.Emojis {
		if emoji.Name == emojiName {
			return emoji
		}
	}

	return nil
}
