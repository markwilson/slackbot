# slackbot

## Usage

``` go
package main

import (
	"os"
	"github.com/nlopes/slack"
	"github.com/markwilson/slackbot"
)

func main() {
	router := slackbot.NewRouter()

	router.AddHandler("echo", []string{"echo"}, EchoHandler{})

	s := slackbot.Slackbot{}
	s.Start(router, os.Getenv("SLACK_TOKEN"))
}

type EchoHandler struct {}

func (h EchoHandler) Handle(rtm *slack.RTM, ev *slack.MessageEvent) {
	rtm.SendMessage(rtm.NewTypingMessage(ev.Channel))

	rtm.SendMessage(rtm.NewOutgoingMessage("Echo: "+ev.Msg.Text, ev.Channel))
}
```
