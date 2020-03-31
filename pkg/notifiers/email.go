package notifiers

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/involvestecnologia/notify/pkg/models"
)

const (
	defaultSMTPServer = "smtp.gmail.com:587"
	messageTemplate   = `From: %s
To: %s
Subject: %s
%s`
)

type gmailNotifier struct {
	auth     smtp.Auth
	defaults models.Options
}

func Gmail(user, password string, opts models.Options) Notifier {
	return &gmailNotifier{auth: smtp.PlainAuth(user, user, password, strings.Split(defaultSMTPServer, ":")[0]),defaults:opts}
}

func (g *gmailNotifier) Notify(from string, to []string, message string, subject string) error {
	m := models.MessageEnvelope{
		From:    from,
		To:      to,
		Message: message,
		Subject: subject,
	}
	m.SetDefaults(g.defaults)
	msg := fmt.Sprintf(messageTemplate, m.From, strings.Join(m.To, ","), m.Subject, m.Message)
	return smtp.SendMail(defaultSMTPServer, g.auth, m.From, m.To, []byte(msg))
}
