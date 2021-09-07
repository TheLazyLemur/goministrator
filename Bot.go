package main

import (
	"fmt"
	"github.com/TheLazyLemur/Goministrator/Voice"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func StartDiscordBot() {
	session, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal(err.Error())
	}

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		log.Fatal(err.Error())
	}

	session.AddHandler(messageCreate)

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	_ = session.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!start-recording" {
		if Voice.Record(s, m) {
			return
		}
	}

	if m.Content == "!end-recording" {
		StopVoice()
	}

	if m.Content == "!create-meeting" {
		println(m.Content)
	}
}

func StopVoice() {
	Voice.Stop()
}
