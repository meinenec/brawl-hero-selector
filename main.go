package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"syscall"
	"os/signal"
	"meinenec/brawl-hero-selector/heroes"
	"strings"
)

func main() {
	token, exists := os.LookupEnv("BOT_TOKEN")
    if !exists {
        panic(fmt.Errorf("BOT_TOKEN not set"))
	}
	
	// Create a new discord session with heroes-bot
	dg, err := discordgo.New("Bot "+ token) 
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Register the brawl func as a callback for "brawl" events
	dg.AddHandler(brawl)

	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func brawl(s *discordgo.Session, m*discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Make m.Content lowercase
	message := strings.ToLower(m.Content)

	// If the message is "brawl" reply with heroes
	if message == "brawl" {
		options := "Pick your Hero!"
		for _, h := range heroes.Assign(3) {
			options = fmt.Sprintf("%s %s %s", options, h.Name, h.Role)
		}
		s.ChannelMessageSend(m.ChannelID, options)
	}

}