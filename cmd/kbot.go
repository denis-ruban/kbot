/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"time"
	"log"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

var (
	teletoken = os.Getenv("TELE_TOKEN")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:   "kbot",
	Aliases: []string{"start"},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kbot started", AppVersion)
		kbot, err:= telebot.NewBot(telebot.Settings{
			URL: "",
			Token: teletoken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable.", err)
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {

			log.Print(m.Message().Payload, m.Text())
			payload := m.Message().Payload

			return err 

			 switch text {
            case "/start", "start":
                return m.Send("👋 Welcome! Type 'help' to see a list of commands")
            case "hello", "hi":
                return m.Send(fmt.Sprintf("Hello I'm Kbot! 🤖 %s", AppVersion))
            case "help":
                return m.Send("Available commands: start, hello, help, version")
            case "version":
                return m.Send(fmt.Sprintf("Kbot version %s", AppVersion))
            default:
                return m.Send("Unknown command. Type 'help' for commands list.")
            }
		})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
