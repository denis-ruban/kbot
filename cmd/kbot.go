package cmd

import (
    "fmt"
    "log"
    "strings"
    "time"
    "os"

    "github.com/spf13/cobra"
    telebot "gopkg.in/telebot.v4"
)

var teletoken = os.Getenv("TELE_TOKEN")

var kbotCmd = &cobra.Command{
    Use:     "kbot",
    Aliases: []string{"start"},
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("kbot started", AppVersion)
        b, err := telebot.NewBot(telebot.Settings{
            Token:  teletoken,
            Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
        })
        if err != nil {
            log.Fatalf("Please check TELE_TOKEN env variable: %v", err)
        }

        b.Handle(telebot.OnText, func(m telebot.Context) error {
            text := strings.TrimSpace(strings.ToLower(m.Text()))
            log.Printf("received text: %q", text)

            if strings.HasPrefix(text, "/") {
                text = strings.TrimPrefix(text, "/")
            }

            switch text {
            case "start":
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

        b.Start()
    },
}

func init() {
    rootCmd.AddCommand(kbotCmd)
}