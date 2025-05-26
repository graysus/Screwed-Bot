package screwedreply

import (
	"fmt"
	"main/botsession"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const SCREWED_APPROVE = "<:screwedapprove:1007517801528950844>"

var screws = []string{
	"screw",
	"mealie",
	"maize",
	"indignan",
	"infuriat",
	"catastrophic",
}

func onMessage(sess *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	lowered := strings.ToLower(m.Content)
	for word := range strings.SplitSeq(lowered, " ") {
		for _, screw := range screws {
			if strings.Contains(word, screw) {
				var message string
				if strings.Contains(word, "@") {
					message = "you think you're so clever huh?"
				} else {
					titlecased := strings.ToUpper(word)[:1] + strings.ToLower(word)[1:]
					message = titlecased + " " + SCREWED_APPROVE
				}
				if _, err := sess.ChannelMessageSendReply(m.ChannelID, message, m.Message.Reference()); err != nil {
					fmt.Println("Error sending message reply in screwedreply: " + err.Error())
				}
				return
			}
		}
	}
}

func Init(sess *botsession.BotSession) {
	sess.S.AddHandler(onMessage)
}
