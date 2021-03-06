package clients

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/involvestecnologia/notify/pkg/models"
	mmModels "github.com/mattermost/mattermost-server/v5/model"
)

const (
	defaultColor = "#9b28e3"
	contentType  = "application/json"
)

type mattermostNotifier struct {
	url      string
	defaults models.Options
	client   *http.Client
}

func NewMattermostNotifier(u string,options models.Options) *mattermostNotifier {
	return &mattermostNotifier{
		url:    u,
		defaults:options,
		client: http.DefaultClient,
	}
}

func (m *mattermostNotifier) Notify(from string, to []string, message string, subject string) error {
	e := models.MessageEnvelope{
		From:    from,
		To:      to,
		Message: message,
		Subject: subject,
	}
	e.SetDefaults(m.defaults)
	return m.notify(e)
}

func (m *mattermostNotifier) CustomNotify(msgs ...models.MattermostMessage) error {
	for _, msg := range msgs {
		err := m.sendMessage([]byte(msg.ToJson()))
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *mattermostNotifier) notify(e models.MessageEnvelope) error {
	for i := range e.To {
		var msg models.MattermostMessage
		msg.Username = e.From
		msg.ChannelName = e.To[i]
		msg.IconURL = "https://avatars0.githubusercontent.com/u/17482172?s=200&v=4"
		msg.Attachments = []*mmModels.SlackAttachment{
			&mmModels.SlackAttachment{
				Color: defaultColor,
				Title: e.Subject,
				Text:  e.Message,
			},
		}
		err := m.sendMessage([]byte(msg.ToJson()))
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *mattermostNotifier) sendMessage(payload []byte) error {
	resp, err := m.client.Post(m.url, contentType, bytes.NewReader(payload))
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
