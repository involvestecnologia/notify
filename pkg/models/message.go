package models

type MessageEnvelope struct {
	From    string
	To      []string
	Message string
	Subject string
}