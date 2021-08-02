package commands

import (
	dg "github.com/bwmarrin/discordgo"
)

// settingsHandler has to determine which subcommand to run within the group
func (c *CommandHandler) settingsHandler(s *dg.Session, i *dg.InteractionCreate) {
	sub1 := i.Data.Options[0]
	switch sub1.Name {
	case "killmail":
		sub2 := sub1.Options[0]
		switch sub2.Name {
		case "channel":
			// ch := sub2.Options[0]
			writeMessage(s, i, "Should have worked")
		default:
			writeMessage(s, i, "I don't know how to do that")
		}
	default:
		writeMessage(s, i, "I don't know how to do that")
	}
}
