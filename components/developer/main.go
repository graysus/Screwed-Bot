package developer

import (
	"fmt"
	"main/botsession"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func onMessage(sess *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "screwed dev ") {
		if m.Author.ID != "292486339692199937" {
			sess.ChannelMessageSendReply(m.ChannelID, "You cannot do this lmaooooo", m.Reference())
			return
		}
		if m.Content == "screwed dev panic" {
			panic("test")
		} else {
			sess.ChannelMessageSendReply(m.ChannelID, "Not a command", m.Reference())
		}
	}
}

func devinter(sess *botsession.BotSession, inter *discordgo.Interaction) {
	if err := sess.RespondWithMessage(inter, "among us"); err != nil {
		fmt.Println("Failure: " + err.Error())
	}
}

func Init(sess *botsession.BotSession) {
	sess.S.AddHandler(onMessage)
	sess.AddAppCommand(devinter, &discordgo.ApplicationCommand{
		Name:        "dev",
		Description: ":programming_forse:",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "command",
				Description: "Command to enter",
				Type:        discordgo.ApplicationCommandOptionString,
			},
		},
	})
}
