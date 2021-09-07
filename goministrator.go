package main

//TODO: Create meeting with participants
//TODO: Send reminder to participants of meeting before it starts
//TODO: Send link to all participants to download meeting
//TODO: Frontend dashboard admin panel and scheduler

import (
	"flag"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"os"
	"os/signal"
	"syscall"
)

var (
	Token    string
	StartBot bool
	StartApi bool
)

func init() {
	flag.StringVar(&Token, "token", "N/A", "The token for your discord bot. Get this token at https://discord.com/developers/applications")
	flag.BoolVar(&StartBot, "bot", false, "Should fire up the discord Bot")
	flag.BoolVar(&StartApi, "api", false, "Should fire up the rest api")
	flag.Parse()
}

func checkForUpcomingMeetings() {
	fmt.Println("This task will run periodically")
}

func executeCronJob() {
	gocron.Every(1).Second().Do(checkForUpcomingMeetings)
	<-gocron.Start()
}

func main() {

	if StartBot {
		go executeCronJob()
		go StartDiscordBot()
	}

	if StartApi {
		go HandleRequests()
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
