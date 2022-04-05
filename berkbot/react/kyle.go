package berkbot_react

import (
	"regexp"

	berkbot_utils "github.com/bsssshhhhhhh/berkbot/berkbot/utils"
	"github.com/bwmarrin/discordgo"
)

func KyleReaction(session *discordgo.Session, message *discordgo.MessageCreate) (*discordgo.Emoji, *string) {
	match, _ := regexp.MatchString("k.*y.*l.*e", message.Content)

	if match {
		emoji := berkbot_utils.FindGuildEmoji(session, message.GuildID, "pink")
		return emoji, nil
	}
	return nil, nil
}
