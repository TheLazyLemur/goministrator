package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/lus/dgc"
	"strings"
)

func HealthCheck(ctx *dgc.Ctx) {
	err := ctx.RespondText("PONG!!!")
	if err != nil {
		panic("The service is NOT healthy: " + err.Error())
	}
}

func CreateMeeting(ctx *dgc.Ctx) {
	var sender = ctx.Event.Author
	var mentions = ctx.Event.Message.Mentions
	var allArgs = make([]string, ctx.Arguments.Amount())
	for i := 0; i < ctx.Arguments.Amount(); i++ {
		allArgs = append(allArgs, ctx.Arguments.Get(i).Raw())
	}

	buildMeetingResponseText(ctx, sender, mentions, allArgs)
}

func InitBotOnServer(ctx *dgc.Ctx) {
	ctx.RespondText("Server " + ctx.Event.GuildID + " has been registered to handle meetings by AdminBot")

	//result, err := ctx.Session.ChannelMessageSend(ctx.Event.ChannelID, "Hello WOrld")
	//if err != nil{
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Println(result)
}

func buildMeetingResponseText(ctx *dgc.Ctx, sender *discordgo.User, mentions []*discordgo.User, allArgs []string) {

	var resultMessage = "Creating a meeting with " + sender.Username + " and "

	for i := 0; i < len(mentions); i++ {
		resultMessage += " " + mentions[i].Username + ", "
	}

	for i := 0; i < len(allArgs); i++ {
		fmt.Println(allArgs[i])
		if strings.Contains(allArgs[i], "date:") {
			var dateString = strings.Split(allArgs[i], ":")

			resultMessage += " On " + dateString[1]
		}
	}

	var err = ctx.RespondText(resultMessage)

	if err != nil {
		fmt.Println("Could not create meeting")
	}
}
