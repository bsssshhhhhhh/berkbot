package berkbot_react

import (
	"log"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func ReeReaction(session *discordgo.Session, message *discordgo.MessageCreate) (*discordgo.Emoji, *string) {
	match, _ := regexp.MatchString("\bre{2,}", message.Content)

	if match {
		guild, err := session.Guild(message.GuildID)
		if err != nil {
			log.Println("Could not get guild " + message.GuildID)
			return nil, nil
		}

		for _, emoji := range guild.Emojis {
			if emoji.Name == "ree" {
				return emoji, nil
			}
		}
	}
	return nil, nil
}
