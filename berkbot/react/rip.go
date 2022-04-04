package berkbot_react

import (
	"log"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func RipReaction(session *discordgo.Session, message *discordgo.MessageCreate) (*discordgo.Emoji, *string) {
	match, _ := regexp.MatchString("r\\.?i\\.?p\\.?", message.Content)

	if match {
		guild, err := session.Guild(message.GuildID)
		if err != nil {
			log.Println("Could not get guild "+message.GuildID+": ", err)
			return nil, nil
		}

		for _, emoji := range guild.Emojis {
			if emoji.Name == "rip" {
				return emoji, nil
			}
		}
	}
	return nil, nil
}
