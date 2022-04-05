package berkbot_react

import (
	"regexp"

	berkbot_utils "github.com/bsssshhhhhhh/berkbot/berkbot/utils"
	"github.com/bwmarrin/discordgo"
)

func ReeReaction(session *discordgo.Session, message *discordgo.MessageCreate) (*discordgo.Emoji, *string) {
	match, _ := regexp.MatchString("\bre{2,}", message.Content)

	if match {
		emoji := berkbot_utils.FindGuildEmoji(session, message.GuildID, "ree")
		return emoji, nil
	}
	return nil, nil
}
