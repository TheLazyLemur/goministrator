package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lus/dgc"

	"database/sql"
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

	sqliteDatabase, _ := sql.Open("sqlite3", "./"+Configuration.DbName) // Open the created SQLite File
	defer sqliteDatabase.Close()                                        // Defer Closing the database

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create(Configuration.DbName) // Create SQLite file

	if err != nil {
		log.Fatal(err.Error())
	}

	file.Close()
	log.Println("sqlite-database.db created")

	dg, err := discordgo.New("Bot " + Configuration.Token)

	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = dg.Open()

	router := dgc.Create(&dgc.Router{
		Prefixes: []string{"-"},
	})

	// Register a simple ping command
	router.RegisterCmd(&dgc.Command{
		Name:        "org",
		Description: "Responds with 'pong!'",
		Usage:       "org",
		Example:     "ping",
		IgnoreCase:  true,
		Handler:     Testing,
	})

	// Initialize the router
	router.Initialize(dg)

	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
