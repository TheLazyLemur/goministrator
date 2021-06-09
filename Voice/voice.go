package Voice

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
	"log"
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

func StartRecording(session *discordgo.Session, guidId string, channelId string) {
	c, err := session.ChannelVoiceJoin(guidId, channelId, true, false)
	if err != nil {
		log.Fatal(err.Error())
	}
	Connection = c
	handleVoice(Connection.OpusRecv)
}

func handleVoice(c chan *discordgo.Packet) {
	files := make(map[uint32]media.Writer)
	for p := range c {
		file, ok := files[p.SSRC]

		if !ok {
			var err error
			file, err = oggwriter.New(fmt.Sprintf("%d.ogg", p.SSRC), 48000, 2)
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
