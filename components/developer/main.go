package developer

import (
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
	sess.S.InteractionRespond(inter, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "why does this take so much code to do",
		},
	})
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
