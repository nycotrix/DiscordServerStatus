package main

import (
	"fmt"

	"github.com/nycotrix/DiscordServerStatus/bot"
	"github.com/nycotrix/DiscordServerStatus/config"
	"github.com/nycotrix/DiscordServerStatus/serverstatus"
)

var version = "0.1"

func init() {
	fmt.Println("Starting Server Status " + version)
}

func main() {
	//read config file
	config.Configure()

	//connect bot to account with token
	bot.Connect(config.Config.Token)

	// add handlers
	bot.AddHandler(serverstatus.MessageHandler)

	//start websocket to listen for messages
	bot.Start()

	//start server status task
	serverstatus.Start()

	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
}
