package main

import (
	"database/sql"
	"github.com/bwmarrin/discordgo"
	"github.com/lus/dgc"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Configuration Config
)

func main() {
	Configuration = ParseConfig()

	createDb()

	dg, err := discordgo.New("Bot " + Configuration.Token)

	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}



	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = dg.Open()

	setupRoutes(dg)

	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	log.Println("Bot is now running.  Press CTRL-C to exit.")

	waitForCloseSignal(dg)
}

func createDb() {
	sqliteDatabase, _ := sql.Open("sqlite3", "./"+Configuration.DbName) // Open the created SQLite File
	defer sqliteDatabase.Close()                                        // Defer Closing the database

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create(Configuration.DbName) // Create SQLite file

	if err != nil {
		log.Fatal(err.Error())
	}

	file.Close()
	log.Println("sqlite-database.db created")
}

func setupRoutes(dg *discordgo.Session) {
	router := dgc.Create(&dgc.Router{
		Prefixes: []string{Configuration.CommandPrefix},
	})
	router.RegisterCmd(&dgc.Command{
		Name:        "ping",
		Description: "Responds with 'pong!'",
		Usage:       "ping",
		Example:     "-ping",
		IgnoreCase:  true,
		Handler:     HealthCheck,
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "init",
		Description: "Make is so the bot will handle meetings, the bot will check every 30 seconds and create the relative channel and roles if a db entry is found",
		Usage:       "init",
		Example:     "-init",
		IgnoreCase:  true,
		Handler:     InitBotOnServer,
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "create",
		Description: "Sets up a meeting between mentioned participants and sender",
		Usage:       "create",
		Example:     "create @user1 @user2 -Date",
		IgnoreCase:  true,
		Flags: []string{
			"greeting",
		},
		Handler:     CreateMeeting,
	})

	router.Initialize(dg)
}


func waitForCloseSignal(dg *discordgo.Session) {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
