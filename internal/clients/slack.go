package clients

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/involvestecnologia/notify/pkg/models"
	mmModels "github.com/mattermost/mattermost-server/v5/model"
)

type slackNotifier struct {
	url    string
	client *http.Client
}

func NewSlackNotifier(u string) *slackNotifier {
	return &slackNotifier{
		url:    u,
		client: http.DefaultClient,
	}
}

func (s *slackNotifier) Notify(from string, to []string, message string, subject string) error {
		msg := models.SlackMessage{
			Attachments: []*mmModels.SlackAttachment{
				&mmModels.SlackAttachment{
					Color: defaultColor,
					Title: subject,
					Text:  message,
				},
			},
		}
	return s.sendMessage([]byte(msg.ToJson()))
}

func (s *slackNotifier) CustomNotify(msgs ...models.SlackMessage) error {
	for _, msg := range msgs {
		err := s.sendMessage([]byte(msg.ToJson()))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *slackNotifier) sendMessage(payload []byte) error {
	resp, err := s.client.Post(s.url, contentType, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("Failed to send MM message %s", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Failed to send MM message with status code %d", resp.StatusCode)
		}
		return fmt.Errorf("Failed to send MM message with status code %d: %s", resp.StatusCode, string(body))
	}
	return nil
}
