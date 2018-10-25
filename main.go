package main

import (
	"fmt"

	"github.com/rafamds/discord-bot/bot"
	"github.com/rafamds/discord-bot/config"
)

func main() {
	err := config.ReadConfiguration()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})

	return
}
