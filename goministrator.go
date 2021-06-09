package main

import (
	"github.com/TheLazyLemur/Goministrator/Voice"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"time"
)

var (
	Token     string
	ChannelID string
	GuildID   string
)

func init() {
	Token = os.Getenv("goministrator-token")
	ChannelID = os.Getenv("goministrator-channelId")
	GuildID = os.Getenv("goministrator-guildId")
}

func main() {
	session, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer func(session *discordgo.Session) {
		err := session.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(session)

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		log.Fatal(err.Error())
	}

	go Voice.StartRecording(session, GuildID, ChannelID)

	time.Sleep(20 * time.Second)
	println("Hello World")
	Voice.Stop()
}
