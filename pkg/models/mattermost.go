package models

import (
	mmModels "github.com/mattermost/mattermost-server/v5/model"
)

type MattermostMessage struct {
	mmModels.IncomingWebhookRequest
}
