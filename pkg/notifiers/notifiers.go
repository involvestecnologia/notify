package notifiers

type Notifier interface {
	Notify(from string, to []string, message string, subject string) error
}

