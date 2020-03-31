package models


type MessageEnvelope struct {
	From    string
	To      []string
	Message string
	Subject string
}

func (m *MessageEnvelope) SetDefaults(defaults Options) {
	if m.From == "" {
		m.From = defaults.DefaultSender
	}
	if len(m.To) == 0 {
		m.To = defaults.DefaultDestinations
	}
	if m.Subject == "" {
		m.Subject = defaults.DefaultSubject
	}
}