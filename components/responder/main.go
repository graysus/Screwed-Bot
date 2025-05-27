package responder

import (
	"fmt"
	"main/botsession"
	"math/rand"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const GARRET_COPYPASTA = "I'm crying in my room rn and almost puked." +
	"I can't accept that garret is fucking gone. I want to kill jumpman25. He is a selfish fat king like fat garret." +
	"Just deciding for no reason to get him removed. He could've given the rights to Grantare and given his shitty self insert to Golden Apple." +
	"But no, he gets to keep him forever to himself. Jumpman25 gets bullied at school and his nuts kicked by the girls..."

const HYDRAULIC_SHREDDER = "The hydraulic-forced feeding tree branch shredder is conducive to reducing the" +
	" volume of fluffy branches and feeding quickly the front pressing roller can prevent the material from " +
	"flowing back and ensure the safety of use the drum cutter rotor structure optimizes the cutting effect " +
	"and can easily crush 15 CM longs to to obtain higher output. The finished product is more suitable for " +
	"making organic fertilizer and ground cover. More things made in China. More tree branch shredder machine" +
	" made in China. If you want to know about any equipment, please let us know in the comment area."

var CUSSES = []string{"FUCK", "SHIT", "BITCH", "ASS", "DICK", "PISS"}

func onMessage(sess *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	lowered := strings.ToLower(m.Content)
	if lowered == "screwed car" {
		r, err := os.Open("assets/images/screwedcar.png")
		if err != nil {
			fmt.Println("Error opening assets/images/screwedcar.png: " + err.Error())
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
			fmt.Println("Error sending message: " + err.Error())
		}
	}
	if strings.Contains(lowered, "garret ") || strings.HasSuffix(lowered, "garret") {
		if _, err := sess.ChannelMessageSendReply(m.ChannelID, GARRET_COPYPASTA, m.Reference()); err != nil {
			fmt.Println("Error sending reply: " + err.Error())
		}
	}
	if strings.Contains(lowered, "hydraulic") || strings.Contains(lowered, "shred") {
		if _, err := sess.ChannelMessageSendReply(m.ChannelID, HYDRAULIC_SHREDDER, m.Reference()); err != nil {
			fmt.Println("Error sending reply: " + err.Error())
		}
	}
	if lowered == "fart" {
		index := rand.Intn(len(CUSSES))
		cuss := CUSSES[index]
		sess.ChannelMessageSendReply(m.ChannelID, cuss, m.Reference())
	}
}

func Init(sess *botsession.BotSession) {
	sess.S.AddHandler(onMessage)
}
