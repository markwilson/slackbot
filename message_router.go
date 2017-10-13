package slackbot

import (
	"github.com/nlopes/slack"
	"fmt"
	"strings"
	"regexp"
)

type Route struct {
	name string
	handler Handler
	patterns []string
}

type Router struct {
	rtm *slack.RTM
	routes map[string]Route
}

func NewRouter() Router {
	var r Router
	r.routes = make(map[string]Route)
	return r
}

// TODO: change patterns to be an array of regex objects instead of string
func (r Router) AddHandler(name string, patterns []string, handler Handler) {
	// TODO: check for a clash with name

	r.routes[name] = Route{
		name: name,
		handler: handler,
		patterns: patterns,
	}
}

func (r Router) Route(ev *slack.MessageEvent) {
	message := strings.ToLower(ev.Msg.Text)

	// TODO: figure out if there's a faster way of handling this route matching
	for name, route := range r.routes {
		for _, pattern := range route.patterns {
			match, _ := regexp.MatchString(pattern, message)

			if match {
				fmt.Printf("Route found: %s\n", name)

				route.handler.Handle(r.rtm, ev)

				return
			}
		}
	}

	handler := NotFoundHandler{}
	handler.Handle(r.rtm, ev)
}

type Handler interface {
	Handle(rtm *slack.RTM, ev *slack.MessageEvent)
}

type NotFoundHandler struct {}

func (h NotFoundHandler) Handle(rtm *slack.RTM, ev *slack.MessageEvent) {
	// TODO: this should only be in a verbose mode
	fmt.Printf("Unhandled: %s\n", ev.Msg.Text)
}
