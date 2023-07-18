package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

const KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot" + Token)
	if err != nil {
		log.Fatal("error creating Discord session, Details:", err)
		return
	}

	// dg.AddHandler()

	dg.Identify.Intents = discordgo.IntentGuildMessages

	// Open websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection, Details:", err)
		return
	}

	log.Println("Bot is now running. Press CTRL + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session
	dg.Close()
}
