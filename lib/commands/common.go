package commands

import (
	dg "github.com/bwmarrin/discordgo"
	"log"
)

func writeMessage(s *dg.Session, i *dg.InteractionCreate, m string) {
	err := s.InteractionRespond(i.Interaction, &dg.InteractionResponse{
		Type: dg.InteractionResponseChannelMessageWithSource,
		Data: &dg.InteractionApplicationCommandResponseData{
			Content: m,
		},
	})
	if err != nil {
		log.Printf("Unable to respond to interaction: %v\n", err)
	}
}
