package notify

import (
	"github.com/slack-go/slack"
)

func NotifySlack(token, channel, message string) error {
	api := slack.New(token)

	_, _, err := api.PostMessage(
		channel,
		slack.MsgOptionText(message, false),
	)
	if err != nil {
		return err
	}

	return nil
}
