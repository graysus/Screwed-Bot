package botsession

import (
	"fmt"

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
	bot.Handlers[cmd.Name] = hndl
	bot.S.ApplicationCommandCreate(bot.S.State.User.ID, "", cmd)
}
