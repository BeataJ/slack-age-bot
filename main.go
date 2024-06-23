package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {
	godotenv.Load()

	tokenbot := os.Getenv("BOT_TOKEN")
	tokenapp := os.Getenv("APP_TOKEN")
	os.Setenv("SLACK_BOT_TOKEN", tokenbot)
	os.Setenv("SLACK_APP_TOKEN", tokenapp)

	bot := slacker.NewClient("SLACK_BOT_TOKEN", os.Getenv("SLACK_APP_TOKEN"))
}
