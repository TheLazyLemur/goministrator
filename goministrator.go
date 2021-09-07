package main

//TODO: Create meeting with participants
//TODO: Send reminder to participants of meeting before it starts
//TODO: Send link to all participants to download meeting
//TODO: Frontend dashboard admin panel and scheduler

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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

func homePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the HomePage!")
	if err != nil {
		return
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":42071", nil))
}

func main() {

	if StartBot {
		go StartDiscordBot()
	}

	if StartApi {
		go handleRequests()
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
