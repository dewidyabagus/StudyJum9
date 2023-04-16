package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/slack-go/slack"
)

/*
	Membuat channel dan generate token slack https://api.slack.com/apps
*/

type slackClient struct {
	channelID string
	*slack.Client
}

func NewSlackClient(channelID, token string) *slackClient {
	return &slackClient{
		channelID: channelID,
		Client:    slack.New(token, slack.OptionDebug(false)),
	}
}

func (s *slackClient) ErrorReporting(ctx context.Context, errMsg, stacktrace string) error {
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Send Message
	attachment := slack.Attachment{
		ServiceName: "Service API Calculation",
		AuthorName:  "Error Reporting",
		Color:       "#eb0729",
		Pretext:     ":boom: Panic Recovery From Service",
		Text:        "There is a problem with the message:\n`" + errMsg + "`\n\nStacktrace:\n```" + stacktrace + "```",
	}
	channel, timestamp, err := s.PostMessageContext(
		ctxWT,
		s.channelID,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		return fmt.Errorf("post message slack: %w", err)
	}
	log.Printf("Message successfully sent to channel %s at %s\n", channel, timestamp)

	return nil
}
