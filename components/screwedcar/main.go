package screwedcar

import (
	"log"
	"main/botsession"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func onMessage(sess *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.ToLower(m.Content) == "screwed car" {
		r, err := os.Open("assets/images/screwedcar.png")
		if err != nil {
			log.Println("Error opening assets/images/screwed.png: " + err.Error())
			return
		}
		defer r.Close()

		MS := &discordgo.MessageSend{
			Reference: m.Message.Reference(),
			File: &discordgo.File{
				Name:   "screwedcar.png",
				Reader: r,
			},
		}

		if _, err := sess.ChannelMessageSendComplex(m.ChannelID, MS); err != nil {
			log.Println("Error sending message: " + err.Error())
		}
	}
}

func Init(sess *botsession.BotSession) {
	sess.S.AddHandler(onMessage)
}
