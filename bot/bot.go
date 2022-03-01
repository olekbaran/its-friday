package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"its-friday/config"
	"net/http"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/robfig/cron.v2"
)

var BotId string
var fridayMessageId []string

type fridayMessageResponse struct {
	Id string
}

func Start() {
	fridayMessageId = make([]string, len(config.FridayMessageChannel))
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
	fmt.Println("Bot is running!")
	fmt.Println("\n-----LOGS-----")

	fridayCron := cron.New()
	fridayCron.AddFunc("0 6 * * FRI", friday)
	fridayCron.Start()

	mondayCron := cron.New()
	mondayCron.AddFunc("0 0 * * MON", monday)
	mondayCron.Start()
}

func friday() {
	for i := 0; i < len(config.FridayMessageChannel); i++ {
		f, err := os.Open("friday-message.json")
		if err != nil {
			fmt.Println("Error with a friday-message.json file!")
		}
		defer f.Close()
		req, err := http.NewRequest("POST", os.ExpandEnv("https://discord.com/api/channels/"+config.FridayMessageChannel[i]+"/messages"), f)
		if err != nil {
			fmt.Println("Error with sending a request")
		}
		req.Header.Set("Authorization", os.ExpandEnv("Bot "+config.Token))
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error with a response")
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)

			if err != nil {
				fmt.Println("Error with a response")
			}

			var fridayMessageResponse1 fridayMessageResponse
			errResponse := json.Unmarshal(bodyBytes, &fridayMessageResponse1)
			if err != nil {
				fmt.Println(errResponse)
			}

			dateFriday := time.Now()
			fmt.Println(dateFriday.Format("01-02-2006") + " Create: " + fridayMessageResponse1.Id)
			fridayMessageId[i] = fridayMessageResponse1.Id
			time.Sleep(time.Second)
		}
	}
}

func monday() {
	for i := 0; i < len(config.FridayMessageChannel); i++ {
		req, err := http.NewRequest("DELETE", "https://discord.com/api/channels/"+config.FridayMessageChannel[i]+"/messages/"+fridayMessageId[i], nil)
		if err != nil {
			fmt.Println("Error with sending a request")
		}
		req.Header.Set("Authorization", "Bot "+config.Token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error with a response")
		}
		defer resp.Body.Close()

		dateMonday := time.Now()
		fmt.Println(dateMonday.Format("01-02-2006") + " Delete: " + fridayMessageId[i])
		time.Sleep(time.Second)
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == config.BotPrefix+"help" || m.Content == config.BotPrefix+" help" {
		help := &discordgo.MessageEmbed{
			Color: 41938,
			Title: "Help 🐬",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "author",
					Value:  "Infos about the developer",
					Inline: true,
				},
				{
					Name:   "ping",
					Value:  "PONG",
					Inline: true,
				},
				{
					Name:   "pong",
					Value:  "PING",
					Inline: true,
				},
				{
					Name:   "when-friday",
					Value:  "Countdown to the next friday",
					Inline: true,
				},
			},
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, help)
	}

	if m.Content == config.BotPrefix+"author" || m.Content == config.BotPrefix+" author" {
		author := &discordgo.MessageEmbed{
			Color: 41938,
			Title: "Author 🐬",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "🧢 Name",
					Value:  "Aleksander Baran",
					Inline: true,
				},
				{
					Name:   "🦋 GitHub",
					Value:  "olek-arsee",
					Inline: true,
				},
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "https://avatars.githubusercontent.com/u/74045117?v=4",
			},
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, author)
	}

	if m.Content == config.BotPrefix+"ping" || m.Content == config.BotPrefix+" ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong 🐬")
	}

	if m.Content == config.BotPrefix+"pong" || m.Content == config.BotPrefix+" pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Ping 🐳")
	}

	if m.Content == config.BotPrefix+"when-friday" || m.Content == config.BotPrefix+" when-friday" {
		today := time.Now().Weekday()
		whenFridayIs := "When?"

		switch time.Friday {
		case today + 0:
			whenFridayIs = "Today 🎉"
		case today + 1:
			whenFridayIs = "Tomorrow ☝🏼"
		case today + 2:
			whenFridayIs = "In two days ✌🏼"
		case today + 3:
			whenFridayIs = "In three days 🎶"
		case today + 4:
			whenFridayIs = "In four days 🍀"
		case today + 5:
			whenFridayIs = "In five days 🖐🏼"
		default:
			fmt.Println("Will be sometime away. ⏱")
		}

		countdown := &discordgo.MessageEmbed{
			Color:       41938,
			Title:       "When friday? 🐬",
			Description: whenFridayIs,
		}

		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, countdown)
	}
}
