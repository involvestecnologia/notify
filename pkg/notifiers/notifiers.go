package notifiers

import (
	"github.com/involvestecnologia/notify/pkg/models"
)

type Notifier interface {
	Notify(e ...models.MessageEnvelope) error
}
