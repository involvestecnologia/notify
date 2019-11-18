package notifiers

import (
	"fmt"
	"os"

	"github.com/involvestecnologia/notify/internal/clients"
	"github.com/involvestecnologia/notify/pkg/models"
)

type mattermostNotifier interface {
	Notifier
	CustomNotify(...models.MattermostMessage) error
}

func MM(url string) mattermostNotifier {
	if url == "" {
		fmt.Println("Missing mattermost webhook URL")
		os.Exit(1)
	}
	return clients.NewMattermostNotifier(url)
}
