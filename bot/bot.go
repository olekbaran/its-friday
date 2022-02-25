package bot

import (
	"fmt"
	"its-friday/config"

	"github.com/bwmarrin/discordgo"
)

var BotId string

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == config.BotPrefix+"ping" || m.Content == config.BotPrefix+" ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
