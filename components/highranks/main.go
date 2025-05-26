package highranks

import (
	"fmt"
	"main/botsession"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func onMessage(sess *discordgo.Session, m *discordgo.MessageCreate) {
	for _, user := range m.Mentions {
		if user.ID != sess.State.User.ID {
			continue
		}
		var content string
		if strings.Contains(strings.ToLower(m.Content), "greg") {
			content = "greg flashing lights warning"
		} else {
			content = "no pinging high ranks"
		}
		if _, err := sess.ChannelMessageSendReply(m.ChannelID, content, m.Reference()); err != nil {
			fmt.Println("Error sending reply in highranks.onMessage: " + err.Error())
			return
		}
	}
}

func Init(sess *botsession.BotSession) {
	sess.S.AddHandler(onMessage)
}
