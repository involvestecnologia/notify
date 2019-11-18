package notifiers

import (
	"fmt"
	"github.com/involvestecnologia/notify/internal/clients"
	"github.com/involvestecnologia/notify/pkg/models"
	"os"
)

type SlackNotifier interface {
	Notifier
	CustomNotify(...models.SlackMessage) error
}

func Slack(url string) SlackNotifier {
	if url == "" {
		fmt.Println("Missing slack webhook URL")
		os.Exit(1)
	}
	return clients.NewSlackNotifier(url)
}
