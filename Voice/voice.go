package Voice

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
	"log"
	"time"
)

var (
	Connection *discordgo.VoiceConnection
)

func createPionRTPPacket(p *discordgo.Packet) *rtp.Packet {
	return &rtp.Packet{
		Header: rtp.Header{
			Version:        2,
			PayloadType:    0x78,
			SequenceNumber: p.Sequence,
			Timestamp:      p.Timestamp,
			SSRC:           p.SSRC,
		},
		Payload: p.Opus,
	}
}

func Record(s *discordgo.Session, m *discordgo.MessageCreate) bool {
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
			go StartRecording(s, vs.GuildID, vs.ChannelID)
			return true
		}
	}
	return false
}

func StartRecording(session *discordgo.Session, guidId string, channelId string) {
	c, err := session.ChannelVoiceJoin(guidId, channelId, true, false)
	if err != nil {
		log.Fatal(err.Error())
	}
	Connection = c
	recordVoice(Connection.OpusRecv)
}

func recordVoice(c chan *discordgo.Packet) {
	files := make(map[uint32]media.Writer)
	now := time.Now()
	for p := range c {
		file, ok := files[p.SSRC]

		if !ok {
			var err error
			file, err = oggwriter.New(fmt.Sprintf("Recordings/%d.ogg", now.Unix()), 48000, 2)
			if err != nil {
				fmt.Printf("failed to create file %d.ogg, giving up on recording: %v\n", p.SSRC, err)
				return
			}
			files[p.SSRC] = file
		}

		newRtp := createPionRTPPacket(p)
		err := file.WriteRTP(newRtp)
		if err != nil {
			fmt.Printf("failed to write to file %d.ogg, giving up on recording: %v\n", p.SSRC, err)
		}
	}
	for _, f := range files {
		err := f.Close()
		if err != nil {
			return
		}
	}
}

func Stop() {
	close(Connection.OpusRecv)
	Connection.Close()
}
