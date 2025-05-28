package botsession

import (
	"fmt"
	"io"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler = func(sess *BotSession, hand *discordgo.Interaction)

type BotSession struct {
	S        *discordgo.Session
	Handlers map[string]CommandHandler
}

func New(sess *discordgo.Session) *BotSession {
	s := &BotSession{
		S:        sess,
		Handlers: map[string]CommandHandler{},
	}
	s.S.AddHandler(func(_ *discordgo.Session, inter *discordgo.InteractionCreate) {
		if inter.Type == discordgo.InteractionApplicationCommand {
			if hand, ok := s.Handlers[inter.ApplicationCommandData().Name]; ok {
				hand(s, inter.Interaction)
			} else {
				s.S.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Command not implemented. Please report this!",
					},
				})
				fmt.Println("Command not implemented: " + inter.ApplicationCommandData().Name + ".")
			}
		}
	})
	return s
}

func (bot *BotSession) Close() error {
	return bot.S.Close()
}

func (bot *BotSession) Load(loaders ...func(s *BotSession)) {
	for _, fun := range loaders {
		fun(bot)
	}
}

func (bot *BotSession) AddAppCommand(hndl CommandHandler, cmd *discordgo.ApplicationCommand) {
	if cmd == nil {
		panic("cmd mustn't be nil")
	}
	bot.Handlers[cmd.Name] = hndl
	if _, err := bot.S.ApplicationCommandCreate(bot.S.State.User.ID, "", cmd); err != nil {
		fmt.Println("Error creating command " + cmd.Name + ": " + err.Error())
	}
}

func (bot *BotSession) RespondWithMessage(inter *discordgo.Interaction, content string) error {
	return bot.S.InteractionRespond(inter, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

func (bot *BotSession) RespondWithMessageAttachment(inter *discordgo.Interaction, name string, rdr io.Reader) error {
	return bot.S.InteractionRespond(inter, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Files: []*discordgo.File{
				{
					Name:   name,
					Reader: rdr,
				},
			},
		},
	})
}
