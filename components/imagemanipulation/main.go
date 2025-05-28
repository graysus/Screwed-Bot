package imagemanipulation

import (
	"fmt"
	"main/botsession"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/gographics/imagick.v3/imagick"
)

const (
	SCREWED_RIGHTHAND_X = 394
	SCREWED_RIGHTHAND_Y = 183
	INTERVALS           = 8
)

func onMessage(bot *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "screwedify" {
		if len(m.Attachments) < 1 {
			bot.ChannelMessageSendReply(m.ChannelID, "Expected attachment", m.Reference())
			return
		}
		atta := m.Attachments[0]
		wand, err := ReadFromHttp(atta.ProxyURL)
		if err != nil {
			fmt.Println("Error reading attachment: " + err.Error())
			return
		}

		wand2 := imagick.NewMagickWand()
		if err := wand2.ReadImage("assets/images/screwed.png"); err != nil {
			fmt.Println(err.Error())
			return
		}

		if err := wand.CompositeImage(wand2, imagick.COMPOSITE_OP_ATOP, true, 0, 0); err != nil {
			fmt.Println("Failed to composite: " + err.Error())
			return
		}

		if rdr, err := WriteWandToReader(wand); err != nil {
			fmt.Println("Error writing to reader: " + err.Error())
		} else if _, err := bot.ChannelFileSend(m.ChannelID, "file.png", rdr); err != nil {
			fmt.Println("Error sending message: " + err.Error())
		}
	}
}

func blah(bot *botsession.BotSession, inter *discordgo.Interaction) {
	switch inter.ApplicationCommandData().Options[0].Name {
	case "bambi_react":
		wand, err := OptionDataToWand(inter, 0)
		if err != nil {
			fmt.Println("Error reading attachment: " + err.Error())
			bot.RespondWithMessage(inter, "There was an error processing your request.")
			return

		}

		wand.ResizeImage(192, 201, imagick.FILTER_CUBIC)

		wand2 := imagick.NewMagickWand()
		if err := wand2.ReadImage("assets/images/livebambireaction.png"); err != nil {
			fmt.Println("Error reading file: " + err.Error())
			bot.RespondWithMessage(inter, "There was an error processing your request.")
			return
		}

		wand2.CompositeImage(wand, imagick.COMPOSITE_OP_ATOP, true, 112, 53)
		// (112, 53)

		if rdr, err := WriteWandToReader(wand2); err != nil {
			fmt.Println("Error writing to reader: " + err.Error())
		} else if err := bot.RespondWithMessageAttachment(inter, "file.png", rdr); err != nil {
			fmt.Println("Error sending message: " + err.Error())
		}

	case "screwed_eat":
		wand, err := OptionDataToWand(inter, 0)
		if err != nil {
			fmt.Println("Error reading attachment: " + err.Error())
			bot.RespondWithMessage(inter, "There was an error processing your request.")
			return
		}

		wandBase := imagick.NewMagickWand()

		if err := wandBase.ReadImage("assets/images/screwed_eat_body.png"); err != nil {
			fmt.Println("Error reading file: " + err.Error())
			bot.RespondWithMessage(inter, "There was an error processing your request.")
			return
		}

		wandOver := imagick.NewMagickWand()

		if err := wandOver.ReadImage("assets/images/screwed_eat_hand.png"); err != nil {
			fmt.Println("Error reading file: " + err.Error())
			bot.RespondWithMessage(inter, "There was an error processing your request.")
			return
		}

		//round(im3.height*(100/im3.width))

		pixwand := imagick.NewPixelWand()
		pixwand.SetColor("#00000000")
		if wand.GetImageHeight() > wand.GetImageWidth() {
			if err := wand.RotateImage(pixwand, 90); err != nil {
				fmt.Println(err.Error())
				bot.RespondWithMessage(inter, "There was an error processing your request.")
				return
			}
		}
		wand.ResizeImage(uint(wand.GetImageWidth()/wand.GetImageHeight()*77), 77, imagick.FILTER_CUBIC)
		if err := wand.RotateImage(pixwand, -120); err != nil {
			fmt.Println(err.Error())
			bot.RespondWithMessage(inter, "There was an error processing your request.")
			return
		}

		wandBase.CompositeImage(wand, imagick.COMPOSITE_OP_OVER, true, SCREWED_RIGHTHAND_X-int(wand.GetImageWidth()/2), SCREWED_RIGHTHAND_Y-int(wand.GetImageHeight()/2))
		wandBase.CompositeImage(wandOver, imagick.COMPOSITE_OP_OVER, true, 0, 0)

		if rdr, err := WriteWandToReader(wandBase); err != nil {
			fmt.Println("Error writing to reader: " + err.Error())
		} else if err := bot.RespondWithMessageAttachment(inter, "file.png", rdr); err != nil {
			fmt.Println("Error sending message: " + err.Error())
		}

	case "screwed_throw":
		wand, err := OptionDataToWand(inter, 0)
		if err != nil {
			fmt.Println("Error reading attachment: " + err.Error())
			bot.RespondWithMessage(inter, "There was an error processing your request.")
			return
		}

		wandTemp := imagick.NewMagickWand()
		if err := wandTemp.ReadImage("assets/images/screwed_throw.png"); err != nil {
			fmt.Println("Error reading file: " + err.Error())
			bot.RespondWithMessage(inter, "There was an error processing your request.")
			return
		}
		wandBase := imagick.NewMagickWand()

		transparent := imagick.NewPixelWand()
		transparent.SetColor("#00000000")

		if err := wandBase.NewImage(1000, 369, transparent); err != nil {
			fmt.Println("Error creating image: " + err.Error())
			bot.RespondWithMessage(inter, "There was an error processing your request.")
			return
		}
		wandBase.CompositeImage(wandTemp, imagick.COMPOSITE_OP_OVER, true, 0, 0)

		if err := wand.ResizeImage(wand.GetImageWidth()/wand.GetImageHeight()*144, 144, imagick.FILTER_CUBIC); err != nil {
			fmt.Println("Error resizing image: " + err.Error())
			bot.RespondWithMessage(inter, "There was an error processing your request.")
			return
		}

		if err := wand.RotateImage(transparent, 15); err != nil {
			fmt.Println("Error rotating image: " + err.Error())
			bot.RespondWithMessage(inter, "There was an error processing your request.")
			return
		}

		curX, curY, curA := 668, 199, 0x20

		curX -= 3 * INTERVALS
		curY -= 1 * INTERVALS

		for i := range INTERVALS {
			curX += 3
			curY += 1
			curA += 128 / INTERVALS
			if i == INTERVALS-1 {
				curA = 0xFF
			}

			x := wand.Clone()
			x.EvaluateImage(imagick.EVAL_OP_MULTIPLY, float64(curA)/255)

			fmt.Println(curA)
			wandBase.CompositeImage(x, imagick.COMPOSITE_OP_OVER, true, curX, curY)
		}

		// intervalX, intervalY = 3, 1

		// stpx, stpy = (708-intervalX*intervals, 249-intervalY*intervals)

		// intervals = 8

		// initialA = 0x20
		// finalA = 0xA0

		if rdr, err := WriteWandToReader(wandBase); err != nil {
			fmt.Println("Error writing to reader: " + err.Error())
		} else if err := bot.RespondWithMessageAttachment(inter, "file.png", rdr); err != nil {
			fmt.Println("Error sending message: " + err.Error())
		}
	}
}

func Init(sess *botsession.BotSession) {
	sess.S.AddHandler(onMessage)
	sess.AddAppCommand(blah, &discordgo.ApplicationCommand{
		Name:        "image",
		Description: "image",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "bambi_react",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "\u26a0\ufe0fLive Bambi Reaction\u26a0\ufe0f",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "shocker",
						Description: "Image for Bambi to react to",
						Type:        discordgo.ApplicationCommandOptionAttachment,
					},
				},
			},
			{
				Name:        "screwed_eat",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Damn my dog hungry :skull:",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "eating",
						Description: "Image for bambi to go fucking devour",
						Type:        discordgo.ApplicationCommandOptionAttachment,
					},
				},
			},
			{
				Name:        "screwed_throw",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "STOP THROWING ME BAMBI!",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "throwing",
						Description: "Image to throw",
						Type:        discordgo.ApplicationCommandOptionAttachment,
					},
				},
			},
		},
	})
}
