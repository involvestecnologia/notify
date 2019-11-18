package notifiers

import (
	"fmt"
	"github.com/involvestecnologia/notify/internal/clients"
	"github.com/involvestecnologia/notify/pkg/models"
	"os"
)

type slackNotifier interface {
	Notifier
	CustomNotify(...models.SlackMessage) error
}

func Slack(url string) slackNotifier {
	if url == "" {
		fmt.Println("Missing slack webhook URL")
		os.Exit(1)
	}
	return clients.NewSlackNotifier(url)
}
