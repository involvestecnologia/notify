package notifiers

import (
	"fmt"
	"os"

	"github.com/involvestecnologia/notify/internal/clients"
	"github.com/involvestecnologia/notify/pkg/models"
)

type MattermostNotifier interface {
	Notifier
	CustomNotify(...models.MattermostMessage) error
}

func MM(url string, opts models.Options) MattermostNotifier {
	if url == "" {
		fmt.Println("Missing mattermost webhook URL")
		os.Exit(1)
	}
	return clients.NewMattermostNotifier(url,opts)
}
