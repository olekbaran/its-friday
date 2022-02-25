package bot

import (
	"fmt"
	"its-friday/config"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/robfig/cron.v2"
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

	c := cron.New()
	c.AddFunc("* * * * *", friday)
	c.Start()
}

func friday() {
	f, err := os.Open("friday-message.json")
	if err != nil {
		fmt.Println("Error with a friday-message.json file!")
	}
	defer f.Close()
	req, err := http.NewRequest("POST", os.ExpandEnv("https://discord.com/api/channels/"+config.FridayMessageChannel+"/messages"), f)
	if err != nil {
		fmt.Println("Error with sending a request")
	}
	req.Header.Set("Authorization", os.ExpandEnv("Bot "+config.Token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error with response")
	}
	defer resp.Body.Close()
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == config.BotPrefix+"ping" || m.Content == config.BotPrefix+" ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong 🏓")
	}
}
