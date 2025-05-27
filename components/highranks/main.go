package highranks

import (
	"fmt"
	"main/botsession"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var MEAN_PHRASES = []string{"shut up", "clam it", "hate you", "fuck you"}

func onMessage(sess *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	hasMeanPhrase := false
	for _, i := range MEAN_PHRASES {
		if strings.Contains(strings.ToLower(m.Content), i) {
			hasMeanPhrase = true
		}
	}
	for _, user := range m.Mentions {
		if user.ID != sess.State.User.ID {
			continue
		}
		var content string
		isGreg := strings.Contains(strings.ToLower(m.Content), "greg")
		if hasMeanPhrase {
			content = "banning you for being mean"
		} else if isGreg {
			content = "greg flashing lights warning"
		} else {
			content = "no pinging high ranks"
		}
		if _, err := sess.ChannelMessageSendReply(m.ChannelID, content, m.Reference()); err != nil {
			fmt.Println("Error sending reply in highranks.onMessage: " + err.Error())
		}
		return
	}
	if strings.Contains(strings.ToLower(m.Content), "screwed") && hasMeanPhrase {
		content := "banning you for being mean"
		if _, err := sess.ChannelMessageSendReply(m.ChannelID, content, m.Reference()); err != nil {
			fmt.Println("Error sending reply in highranks.onMessage: " + err.Error())
		}
	}
}

func Init(sess *botsession.BotSession) {
	sess.S.AddHandler(onMessage)
}
