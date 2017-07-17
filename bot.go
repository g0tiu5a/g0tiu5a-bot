package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/g0tiu5a/ctftime"
	"github.com/nlopes/slack"
)

const (
	botName = "保登 心愛"
	botIcon = ":cocoa-1:"
)

type Bot struct {
	api *slack.Client
	rtm *slack.RTM
}

func (bot *Bot) handleResponse(ev *slack.MessageEvent) error {
	command := strings.Fields(ev.Text)

	if len(command) <= 0 {
		return nil
	}

	switch command[0] {
	case "ping":
		bot.rtm.SendMessage(bot.rtm.NewOutgoingMessage("pong", ev.Channel))
		return nil

	case "ctftime":
		err := bot.ctftime(command, ev.Channel)
		return err
	}
	return nil
}

func (bot *Bot) ctftime(commands []string, channel string) error {
	if len(commands) != 2 {
		return nil
	}

	var attachments []slack.Attachment
	const timeLayout = "2006/01/02 15:04:05 UTC"

	switch commands[1] {
	case "event":
		events := ctftime.GetAPIData()
		for _, event := range events {
			attachment := slack.Attachment{
				Color:     "#F35A00",
				Title:     event.Title,
				TitleLink: event.Url,
				Fields: []slack.AttachmentField{
					{
						Title: "format",
						Value: event.Format,
						Short: true,
					},
					{
						Title: "weight",
						Value: fmt.Sprintf("%f", event.Weight),
						Short: true,
					},
					{
						Title: "start",
						Value: event.Start.Format(timeLayout),
						Short: true,
					},
					{
						Title: "finish",
						Value: event.Finish.Format(timeLayout),
						Short: true,
					},
				},
			}

			attachments = append(attachments, attachment)
		}
	}

	params := slack.PostMessageParameters{
		Attachments: attachments,
		Username:    botName,
		IconEmoji:   botIcon,
	}

	_, _, err := bot.api.PostMessage(channel, "", params)
	return err
}

func (bot *Bot) run() int {
	rtm := bot.rtm
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				err := bot.handleResponse(ev)
				if err != nil {
					log.Fatalf("Error: %v\n", ev)
					bot.rtm.SendMessage(bot.rtm.NewOutgoingMessage(fmt.Sprintf("Something went wrong... Your input is %s", ev.Text), ev.Channel))
				}

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1
			}
		}
	}
}

func main() {
	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
		log.Fatalf("Set SLACK_API_TOKEN in your environment\n")
		os.Exit(1)
	}

	bot := Bot{}
	bot.api = slack.New(token)
	bot.rtm = bot.api.NewRTM()
	os.Exit(bot.run())
}
