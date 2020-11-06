package main

import (
    "database/sql"
    "fmt"
    "github.com/bwmarrin/discordgo"
    "github.com/lus/dgc"
    "log"
    "strings"
    "time"
)

func Testing(ctx *dgc.Ctx)  {
    HandleCommands(ctx.Session, ctx.Event)
}

func HandleCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.Author.ID == s.State.User.ID {
        return
    }

    if !strings.HasPrefix(m.Content, Configuration.CommandPrefix) {
        return
    }

    if strings.Contains(m.Content, "meeting-create") {

        sqliteDatabase, _ := sql.Open("sqlite3", "./"+Configuration.DbName) // Open the created SQLite File
        defer sqliteDatabase.Close()                                        // Defer Closing the database

        var createMeetingTableSQL = `CREATE TABLE meeting (
        "idMeeting" integer NOT NULL PRIMARY KEY AUTOINCREMENT,     
        "idOrganiser" TEXT,
        "idSecond" TEXT,
        "date" TEXT
      );`

        statement, err := sqliteDatabase.Prepare(createMeetingTableSQL) // Prepare SQL Statement

        if err != nil {
            log.Fatal(err.Error())
        }
        statement.Exec()

        dateTime := time.Now()
        currentTime, month, day := time.Now().Date()

        var response = fmt.Sprintf(`INSERT INTO meeting(idOrganiser, idSecond, date) VALUES ("%s", "%s", "%s")`, m.Author.ID, m.Mentions[0].ID, dateTime)

        s.ChannelMessageSend(m.ChannelID,fmt.Sprintf("%s, %s have a meeting organised for %d %s %d", m.Author.Mention(), m.Mentions[0].Mention(), currentTime, month, day))

        statement2, err2 := sqliteDatabase.Prepare(response) // Prepare statement.
        statement2.Exec()

        log.Println(statement2)
        log.Println(err2)
    }
}
