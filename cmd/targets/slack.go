package targets

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"isnan.eu/alerting/cmd/models"
)

var slackToken string = os.Getenv("SLACK_TOKEN")

var channelId string = os.Getenv("SLACK_CHANNEL_ID")

type (
	slackNotifier struct {
		name        string
		slackClient *slack.Client
	}
)

func (n slackNotifier) GetName() string {
	return n.name
}

func (n slackNotifier) SendAlert(alert *models.AlertMessage) error {
	content := string(alert.Content)
	if content != "" {

		attachment := slack.Attachment{
			Pretext: alert.SourceDescription,
			Text:    content,
			/*
				// Color Styles the Text, making it possible to have like Warnings etc.
				Color: "#36a64f",
				// Fields are Optional extra data!
				Fields: []slack.AttachmentField{
					{
						Title: "Date",
						Value: time.Now().String(),
					},
				},
			*/
		}

		_, _, err := n.slackClient.PostMessage(channelId, slack.MsgOptionAttachments(attachment))

		if err != nil {
			log.Errorf("Failed to send alert to %s", n.name)
			return err
		}

	}
	return nil
}

func NewSlackTarget() Target {
	target := slackNotifier{
		name:        "slack",
		slackClient: slack.New(slackToken),
	}
	return target
}
