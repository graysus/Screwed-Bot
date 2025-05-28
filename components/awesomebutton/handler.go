package awesomebutton

import (
	"bytes"
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

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
