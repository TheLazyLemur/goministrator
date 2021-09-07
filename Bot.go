package main

import (
	"fmt"
	"github.com/TheLazyLemur/Goministrator/Voice"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
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

	if strings.HasPrefix(m.Content, "!help") {
		_, err := s.ChannelMessageSend(m.ChannelID, "```!start-recording``` The bot will join your room and record audio\n```!end-recording``` Will cause the bot to stop recording audio and eventually leve the room")
		if err != nil {
			log.Fatalf("Could not send message due to: %s", err.Error())
		}
	}

	if strings.HasPrefix(m.Content, "!start-recording") {
		if Voice.Record(s, m) {
			return
		}
	}

	if strings.HasPrefix(m.Content, "!end-recording") {
		Voice.Stop()
	}

	if strings.HasPrefix(m.Content, "!create-meeting") {
		stringSlice := strings.Split(m.Content, "!create-meeting")
		content := strings.TrimSpace(stringSlice[1])
		participantIds := strings.Split(content, ",")

		var participants []*discordgo.User

		for _, element := range participantIds {
			usr, _ := s.User(element)
			participants = append(participants, usr)
		}

		for _, participant := range participants {
			if participant.Bot {
				continue
			}
			userChn, err := s.UserChannelCreate(participant.ID)
			if err != nil {
				log.Fatalf("Could not create user channel: %v", err.Error())
			}

			_, err = s.ChannelMessageSend(userChn.ID, "You have been added to a meeting at this time")
			if err != nil {
				log.Fatalf("Could not send message to user channel: %v", err.Error())
			}
		}
	}
}
