package randomscrewed

import (
	"log"
	"main/botsession"
	"math/big"
	"os"

	"github.com/bwmarrin/discordgo"
)

func onMessage(sess *discordgo.Session, m *discordgo.MessageCreate) {
	screwyHash := BytesHash([]byte(m.Content))
	screwyHash.Sub(screwyHash, big.NewInt(28)).Mod(screwyHash, big.NewInt(150))
	if screwyHash.Cmp(big.NewInt(0)) == 0 {
		r, err := os.Open("assets/images/screwed.png")
		if err != nil {
			log.Println("Error opening assets/images/screwed.png: " + err.Error())
			return
		}
		defer r.Close()

		MS := &discordgo.MessageSend{
			Reference: m.Message.Reference(),
			File: &discordgo.File{
				Name:   "screwed.png",
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
