package randomscrewed

import (
	"log"
	"main/botsession"
	"math/big"
	"os"

	"github.com/bwmarrin/discordgo"
)

func checkScrewed(content string) bool {
	screwyHash := BytesHash([]byte(content))
	screwyHash.Sub(screwyHash, big.NewInt(28)).Mod(screwyHash, big.NewInt(150))
	return screwyHash.Cmp(big.NewInt(0)) == 0
}

func onMessage(sess *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	if checkScrewed(m.Content) {
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
	sess.AddAppCommand(screwedify, &discordgo.ApplicationCommand{
		Name:        "screwedify",
		Description: "Screwed-ify a message...",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "base",
				Description: "The base message content",
				Type:        discordgo.ApplicationCommandOptionString,
			},
		},
	})
}
