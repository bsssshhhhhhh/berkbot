package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bsssshhhhhhh/berkbot/berkbot"
	"github.com/bwmarrin/discordgo"
)

var (
	/* API Token */
	Token string
)

/* Setup command line flags */
func Init() {
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.Parse()

	if Token == "" {
		Token = os.Getenv("DISCORD_API_KEY")
	}
}

func main() {
	Init()

	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	err = discord.Open()
	if err != nil {
		fmt.Println("Could not connect to Discord: ", err)
		return
	}

	berkbot.AddCommands(discord)
}