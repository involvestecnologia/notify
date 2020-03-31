package main

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/involvestecnologia/notify/internal/config"
	"github.com/involvestecnologia/notify/pkg/models"
	"github.com/involvestecnologia/notify/pkg/notifiers"
)

var (
	conf    *viper.Viper
	from    string
	to      string
	subject string
	msg     string
)

func init() {
	conf = config.Load()
}

func main() {
	var cmdMM = &cobra.Command{
		Use:   "mm -t [TARGETS] -m MESSAGE [-f -s]",
		Short: "Sends a message to mattermost",
		Long:  `Sends a message to a mattermost channel`,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			e := models.MessageEnvelope{
				From:    from,
				To:      strings.Split(to, ","),
				Subject: subject,
				Message: msg,
			}
			sendMM(e)
		},
	}

	var cmdSlack = &cobra.Command{
		Use:   "slack [-s -m]",
		Short: "Sends a message to slack",
		Long:  `Sends a message to a slack channel`,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			e := models.MessageEnvelope{
				Subject: subject,
				Message: msg,
			}
			sendSlack(e)
		},
	}

	var rootCmd = &cobra.Command{Use: "notify"}

	cmdMM.Flags().StringVarP(&from, "from", "f", "", "The name from the sender")
	cmdMM.Flags().StringVarP(&to, "to", "t", "", "The destination channel or user")
	cmdMM.Flags().StringVarP(&subject, "subject", "s", "", "The subject of the message")
	cmdMM.Flags().StringVarP(&msg, "message", "m", "", "The message body")
	cmdMM.MarkFlagRequired("to")
	cmdMM.MarkFlagRequired("message")

	cmdSlack.Flags().StringVarP(&subject, "subject", "s", "", "The subject of the message")
	cmdSlack.Flags().StringVarP(&msg, "message", "m", "", "The message body")
	cmdSlack.MarkFlagRequired("message")

	rootCmd.AddCommand(cmdMM, cmdSlack)

	rootCmd.Execute()
}

func sendMM(e models.MessageEnvelope) {
	notifiers.MM(conf.GetString("webhook.mm"),models.Options{}).Notify(e.From,e.To,e.Message,e.Subject)
}

func sendSlack(e models.MessageEnvelope) {
	notifiers.Slack(conf.GetString("webhook.slack")).Notify("",nil,e.Message,e.Subject)
}
