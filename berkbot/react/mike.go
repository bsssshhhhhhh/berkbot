package berkbot_react

import (
	"regexp"

	berkbot_utils "github.com/bsssshhhhhhh/berkbot/berkbot/utils"
	"github.com/bwmarrin/discordgo"
)

func MikeReaction(session *discordgo.Session, message *discordgo.MessageCreate) (*discordgo.Emoji, *string) {
	match, _ := regexp.MatchString("mike", message.Content)

	if match {
		emoji := berkbot_utils.FindGuildEmoji(session, message.GuildID, "mike")
		return emoji, nil
	}
	return nil, nil
}
