package models

import (
	mmModels "github.com/mattermost/mattermost-server/model"
)

type MattermostMessage struct {
	mmModels.IncomingWebhookRequest
}
