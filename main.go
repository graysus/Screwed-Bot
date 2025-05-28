package main

import (
	"fmt"
	"log"
	"main/botsession"
	"main/components/amam"
	"main/components/awesomebutton"
	"main/components/developer"
	"main/components/highranks"
	"main/components/imagemanipulation"
	"main/components/randomscrewed"
	"main/components/responder"
	"main/components/screwedreply"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	var strtoken string
	if rawtoken, err := os.ReadFile("TOKEN"); err != nil {
		log.Fatalln("Could not read token file: " + err.Error())
	} else {
		strtoken = string(rawtoken)
	}

	sess, err := discordgo.New("Bot " + strtoken)
	if err != nil {
		log.Fatalln("Error creating session: " + err.Error())
	}
	sess.Identify.Intents = discordgo.IntentsAll
	if err := sess.Open(); err != nil {
		log.Fatalln("Error opening session: " + err.Error())
	}
	bot := botsession.New(sess)
	defer func() {
		if err := bot.Close(); err != nil {
			fmt.Println("Error while closing session: " + err.Error())
		}
	}()

	Init(bot)

	x := make(chan os.Signal, 1)
	signal.Notify(x, syscall.SIGINT, syscall.SIGTERM)
	<-x
}

func Init(sess *botsession.BotSession) {
	sess.Load(
		randomscrewed.Init,
		screwedreply.Init,
		developer.Init,
		amam.Init,
		highranks.Init,
		responder.Init,
		awesomebutton.Init,
		imagemanipulation.Init)
}
