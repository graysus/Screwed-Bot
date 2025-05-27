package awesomebutton

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/botsession"
	"slices"
	"strconv"
	"strings"

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

var deadUsersByID = map[string][]string{}

func interhand(bot *discordgo.Session, inter *discordgo.InteractionCreate) {
	if inter.Type == discordgo.InteractionMessageComponent {
		dat := inter.MessageComponentData()
		if len(dat.CustomID) < 4 || dat.CustomID[:4] != "Json" {
			return
		}

		jsondat := dat.CustomID[4:]
		dec := json.NewDecoder(strings.NewReader(jsondat))

		jsonmap := map[string]any{}

		if err := dec.Decode(&jsonmap); err != nil {
			return
		}

		if val, ok := jsonmap["type"]; ok {
			switch val {
			case "AwesomeButton":
				// assume everything is already filled out (bad idea)
				count := jsonmap["count"].(float64)
				count += 1
				jsonmap["count"] = count

				stringbuf := bytes.NewBufferString("")

				enc := json.NewEncoder(stringbuf)
				enc.Encode(jsonmap)

				row := inter.Message.Components[0].(*discordgo.ActionsRow)
				btn := row.Components[0].(*discordgo.Button)
				btn.CustomID = "Json" + stringbuf.String()
				btn.Label = "Awesome fucking button (" + strconv.Itoa(int(count)) + ")"

				if err := bot.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseUpdateMessage,
					Data: &discordgo.InteractionResponseData{
						Content:    inter.Message.Content,
						Components: []discordgo.MessageComponent{row},
					},
				}); err != nil {
					fmt.Println("uhhh..." + err.Error())
				}
			case "DeathButton":
				// assume everything is already filled out (bad idea)
				user := inter.User
				if user == nil {
					user = inter.Member.User
				}
				val, ok := deadUsersByID[inter.Message.ID]
				if !ok {
					val = []string{}
				}
				alreadyDied := slices.Contains(val, user.ID)
				if alreadyDied {
					if err := bot.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "You're already dead!",
							Flags:   1 << 6,
						},
					}); err != nil {
						fmt.Println("Failed to respond: " + err.Error())
					}
				} else {
					val = append(val, user.ID)
					deadUsersByID[inter.Message.ID] = val
					if err := bot.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "<@!" + user.ID + "> has died!",
						},
					}); err != nil {
						fmt.Println("Failed to respond: " + err.Error())
					}
				}
			case "respond":
				msg := jsonmap["with"].(string)
				if err := bot.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: msg,
						Flags:   1 << 6,
					},
				}); err != nil {
					fmt.Println("Failed to respond: " + err.Error())
				}
			}

		}
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
