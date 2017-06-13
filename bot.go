package main

import (
	"log"
	"os"
	"regexp"

	"github.com/nlopes/slack"
)

func messageEvent(ev *slack.MessageEvent) (string, error) {
	hearing := ev.Text
	ret, _ := regexp.MatchString(`^ping`, hearing)
	if ret {
		return "ping", nil
	}

	return "", nil
}

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				ret, err := messageEvent(ev)
				if err != nil {
					log.Println(err)
				}
				rtm.SendMessage(rtm.NewOutgoingMessage(ret, ev.Channel))

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

	api := slack.New(token)
	os.Exit(run(api))
}
