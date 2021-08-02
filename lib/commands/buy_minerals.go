package commands

import (
	"fmt"
	dg "github.com/bwmarrin/discordgo"
)

func (c *CommandHandler) buyMineralsHandler(s *dg.Session, i *dg.InteractionCreate) {
	order := map[string]int64{}

	for _, v := range i.Data.Options {
		order[v.Name] = v.IntValue()
	}

	fmt.Printf("\n%+v\n\n", order)
	writeMessage(s, i, "Order placed")
}
