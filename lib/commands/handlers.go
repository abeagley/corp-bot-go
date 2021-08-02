package commands

import (
	"cloud.google.com/go/firestore"
	dg "github.com/bwmarrin/discordgo"
)

type CommandHandlers map[string]func(s *dg.Session, i *dg.InteractionCreate)

type CommandHandler struct {
	Db *firestore.Client
}

func (c *CommandHandler) GetMap() CommandHandlers {
	return CommandHandlers{
		"buy-min": c.buyMineralsHandler,
		"settings": c.settingsHandler,
	}
}
