package main

import (
    "fmt"
    "os"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/tarm/serial"
    "time"
)

// The first argument must be the Arduino serial communication port,
// the second one, the Telegram API
func main() {
    // Arduino setup
    config := &serial.Config {
        Name: os.Args[1],
        Baud: 9600,
        ReadTimeout: 2 * time.Second,
    }
    arduino, oops := serial.OpenPort(config)
    defer arduino.Close()
    if oops != nil {
        fmt.Println("Couldn't open Arduino")
        return
    }

    // Telegram bot
    bot, err := telegram.NewBotAPI(os.Args[2])
	if err != nil {
        panic(err)
	}
	u := telegram.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
    fmt.Println("--- # Welcome to Autohome!")
	for update := range updates {
		if update.Message != nil {
            arduino.Write([]byte(update.Message.Text + "\n"))
            buffer := make([]byte, 128)
            arduino.Read(buffer)
            answer := string(buffer)
            msg := telegram.NewMessage(update.Message.Chat.ID, answer)
            bot.Send(msg)
		}
	}
}
