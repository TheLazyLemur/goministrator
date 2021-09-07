package main

import (
	"encoding/binary"
	"fmt"
	"github.com/TheLazyLemur/Goministrator/Voice"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Token  string
	buffer [][]byte
)

func init() {
	Token = "ODUxNDU0MTg1NzYwMjI3MzQ4.YL4ggQ.Jh3an4JdLRUSQZi-PT6bGM7wUZ8"
	buffer = make([][]byte, 0)
}

func main() {
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
		if record(s, m) {
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

func record(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return true
	}

	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		return true
	}

	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			err := playSound(s, vs.GuildID, vs.ChannelID)
			if err != nil {
				return true
			}
			go Voice.StartRecording(s, vs.GuildID, vs.ChannelID)
			return true
		}
	}
	return false
}

// loadSound attempts to load an encoded sound file from disk.
func loadSound(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return err
	}

	var opusLen int16

	for {
		err = binary.Read(file, binary.LittleEndian, &opusLen)

		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return err
			}
			return nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		InBuf := make([]byte, opusLen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		buffer = append(buffer, InBuf)
	}
}

// playSound plays the current buffer to the provided channel.
func playSound(s *discordgo.Session, guildID, channelID string) (err error) {
	err = loadSound("airhorn.dca")
	if err != nil {
		println(err.Error())
	}

	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	time.Sleep(250 * time.Millisecond)

	err = vc.Speaking(true)
	if err != nil {
		return err
	}

	for _, buff := range buffer {
		vc.OpusSend <- buff
	}

	err = vc.Speaking(false)
	if err != nil {
		return err
	}

	time.Sleep(250 * time.Millisecond)

	err = vc.Disconnect()
	if err != nil {
		return err
	}

	return nil
}
