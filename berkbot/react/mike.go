package berkbot_react

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func MikeReaction(session *discordgo.Session, message *discordgo.MessageCreate) (*discordgo.Emoji, *string) {
	match, _ := regexp.MatchString("mike", message.Content)

	if match {
		guild, err := session.Guild(message.GuildID)
		if err != nil {
			fmt.Println("Could not get guild "+message.GuildID+": ", err)
			return nil, nil
		}

		for _, emoji := range guild.Emojis {
			if emoji.Name == "mike" {
				return emoji, nil
			}
		}
	}
	return nil, nil
}
