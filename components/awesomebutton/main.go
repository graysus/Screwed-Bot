package awesomebutton

import (
	"fmt"
	"main/botsession"

	"github.com/bwmarrin/discordgo"
)

func awesomefuckingbutton(bot *botsession.BotSession, inter *discordgo.Interaction) {
	if err := bot.S.InteractionRespond(inter, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Look at this awesome fucking button!",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Awesome fucking button (0)",
							Style:    discordgo.PrimaryButton,
							CustomID: "Json{\"type\": \"AwesomeButton\", \"count\": 0}",
						},
						discordgo.Button{
							Label:    "<<<< OH MY GOD ITS THE AWESOME FUCKING BUTTON",
							Style:    discordgo.SecondaryButton,
							CustomID: "Json{\"type\": \"respond\", \"with\": \"Click on the awesome fucking button!\"}",
						},
					},
				},
			},
		},
	}); err != nil {
		fmt.Println("Error sending stuff: " + err.Error())
	}
}

func deathbutton(bot *botsession.BotSession, inter *discordgo.Interaction) {
	if err := bot.S.InteractionRespond(inter, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Look at this DEATH BUTTON......",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "DIE......",
							Style:    discordgo.DangerButton,
							CustomID: "Json{\"type\": \"DeathButton\"}",
						},
						discordgo.Button{
							Label:    "<<<< OH MY GOD ITS THE DEATH BUTTON",
							Style:    discordgo.SecondaryButton,
							CustomID: "Json{\"type\": \"respond\", \"with\": \"Click on the death button! Or don't, probably...\"}",
						},
					},
				},
			},
		},
	}); err != nil {
		fmt.Println("Error sending stuff: " + err.Error())
	}
}

func Init(sess *botsession.BotSession) {
	sess.AddAppCommand(awesomefuckingbutton, &discordgo.ApplicationCommand{
		Name:        "awesomefuckingbutton",
		Description: "AWESOME FUCKING BUTTON",
	})
	sess.AddAppCommand(deathbutton, &discordgo.ApplicationCommand{
		Name:        "deathbutton",
		Description: "i'm dead :skull: :skull_crossbones:",
	})
	sess.S.AddHandler(interhand)
}
