package berkbot_react

import (
	"regexp"

	berkbot_utils "github.com/bsssshhhhhhh/berkbot/berkbot/utils"
	"github.com/bwmarrin/discordgo"
)

func RipReaction(session *discordgo.Session, message *discordgo.MessageCreate) (*discordgo.Emoji, *string) {
	match, _ := regexp.MatchString("r\\.?i\\.?p\\.?", message.Content)

	if match {
		emoji := berkbot_utils.FindGuildEmoji(session, message.GuildID, "rip")
		return emoji, nil
	}
	return nil, nil
}
