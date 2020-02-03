package models

import (
	"encoding/json"
	mmModels "github.com/mattermost/mattermost-server/v5/model"
)

func (s *SlackMessage) ToJson() string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

type SlackMessage struct {
	SlackText   string                      `json:"text"`
	Blocks      []SlackBlock                `json:"blocks"`
	Attachments []*mmModels.SlackAttachment `json:"attachments"`
}

type SlackBlock struct {
	Type      string          `json:"type"`
	Text      *SlackText      `json:"text,omitempty"`
	BlockID   *string         `json:"block_id,omitempty"`
	Accessory *SlackAccessory `json:"accessory,omitempty"`
	Fields    []SlackText     `json:"fields"`
}

type SlackAccessory struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

type SlackText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
