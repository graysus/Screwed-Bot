package amam

import (
	"fmt"
	"main/botsession"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type AmAmType struct {
	Aliases []string
	ID      string
}

var AMAMTYPES = []AmAmType{
	{
		Aliases: []string{"merry am am", "am am christmas", "christmas am am", "holly jolly am am", "santa am am", "festive am am"},
		ID:      "merryamam:1029169615412138016",
	},
	{
		Aliases: []string{"festival am am", "bam bam", "bam am", "rano am am"},
		ID:      "bamam:1028385639265218712",
	},
	{
		Aliases: []string{"expung am am", "expunged am am", "unfairness am am", "unfairness expunged am am", "unfairness bambi am am", "unfair am am", "unfair expunged am am", "unfair bambi am am"},
		ID:      "unfairamam:1029184079893119056",
	},
	{
		Aliases: []string{"goofy am am", "goofy ahh am am", "silly am am", "am amn't", "am amnt", "fake am am", "walmart am am", "am am 2"},
		ID:      "goofy_ahh_am_am:1029186104533987409",
	},
	{
		Aliases: []string{"am am"},
		ID:      "amam:1023759264512213033",
	},
}

/*
 */

func onMessage(sess *discordgo.Session, m *discordgo.MessageCreate) {
	lowered := strings.ToLower(m.Content)
	for _, amam := range AMAMTYPES {
		for _, alias := range amam.Aliases {
			if strings.Contains(lowered, alias) {
				if err := sess.MessageReactionAdd(m.ChannelID, m.ID, amam.ID); err != nil {
					fmt.Println("Error reacting: " + err.Error())
				}
				return
			}
		}
	}
}

func Init(sess *botsession.BotSession) {
	sess.S.AddHandler(onMessage)
}
