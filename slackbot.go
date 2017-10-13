package slackbot

import (
	"fmt"
	"log"
	"os"
	"github.com/nlopes/slack"
)

type Slackbot struct {}

func (s Slackbot) Start(router Router, slack_token string) {
	api := slack.New(slack_token)
	logger := log.New(os.Stdout, "stage-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)

	rtm := api.NewRTM()
	router.rtm = rtm
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			router.Route(ev)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return
		}
	}
}

