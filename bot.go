package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

type Bot struct {
	api *slack.Client
	rtm *slack.RTM
}

func (bot *Bot) handleResponse(ev *slack.MessageEvent) error {
	command := strings.Fields(ev.Text)

	switch command[0] {
	case "ping":
		bot.rtm.SendMessage(bot.rtm.NewOutgoingMessage("pong", ev.Channel))
		return nil
	}
	return nil
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
