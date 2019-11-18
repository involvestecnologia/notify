package notifiers

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/involvestecnologia/notify/pkg/models"
)

const (
	defaultSMTPServer = "stmp.gmail.com:587"
	messageTemplate   = `From: %s
To: %s
Subject: %s
%s`
)

type gmailNotifier struct {
	auth smtp.Auth
}

func Gmail(user, password string) Notifier {
	return &gmailNotifier{auth: smtp.PlainAuth(user, user, password, strings.Split(defaultSMTPServer, ":")[0])}
}

func (g *gmailNotifier) Notify(messages ...models.MessageEnvelope) error {
	for _, m := range messages {
		msg := fmt.Sprintf(messageTemplate, m.From, strings.Join(m.To, ","), m.Subject, m.Message)
		if err := smtp.SendMail(defaultSMTPServer, g.auth, m.From, m.To, []byte(msg)); err != nil {
			return err
		}
	}
	return nil
}
