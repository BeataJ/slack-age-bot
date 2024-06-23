package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Commant Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	godotenv.Load()

	tokenbot := os.Getenv("BOT_TOKEN")
	tokenapp := os.Getenv("APP_TOKEN")
	os.Setenv("SLACK_BOT_TOKEN", tokenbot)
	os.Setenv("SLACK_APP_TOKEN", tokenapp)

	bot := slacker.NewClient("SLACK_BOT_TOKEN", os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
}
