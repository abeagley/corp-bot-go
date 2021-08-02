package commands

import dg "github.com/bwmarrin/discordgo"

var Commands = []*dg.ApplicationCommand{
	{
		Name: "buy-min",
		Description: "Purchase minerals through the corp",
		Options: []*dg.ApplicationCommandOption{
			{
				Name: "tritanium",
				Description: "Amount of tritanium",
				Required: false,
				Type: dg.ApplicationCommandOptionInteger,
			},
			{
				Name: "pyerite",
				Description: "Amount of pyerite",
				Required: false,
				Type: dg.ApplicationCommandOptionInteger,
			},
			{
				Name: "mexallon",
				Description: "Amount of mexallon",
				Required: false,
				Type: dg.ApplicationCommandOptionInteger,
			},
			{
				Name: "isogen",
				Description: "Amount of isogen",
				Required: false,
				Type: dg.ApplicationCommandOptionInteger,
			},
			{
				Name: "nocxium",
				Description: "Amount of nocxium",
				Required: false,
				Type: dg.ApplicationCommandOptionInteger,
			},
			{
				Name: "zydrine",
				Description: "Amount of zydrine",
				Required: false,
				Type: dg.ApplicationCommandOptionInteger,
			},
			{
				Name: "megacyte",
				Description: "Amount of megacyte",
				Required: false,
				Type: dg.ApplicationCommandOptionInteger,
			},
			{
				Name: "morphite",
				Description: "Amount of morphite",
				Required: false,
				Type: dg.ApplicationCommandOptionInteger,
			},
		},
	},
	{
		Name:        "settings",
		Description: "Configure your Corp Bot's settings",
		Options: []*dg.ApplicationCommandOption{
			{
				Name:        "admins",
				Description: "Roles who can control settings",
				Type:        dg.ApplicationCommandOptionSubCommandGroup,
				Options: []*dg.ApplicationCommandOption{
					{
						Name:        "add",
						Description: "Adds a role to control settings",
						Type:        dg.ApplicationCommandOptionSubCommand,
						Options: []*dg.ApplicationCommandOption{
							{
								Name:        "role",
								Description: "The role to add",
								Required:    true,
								Type:        dg.ApplicationCommandOptionRole,
							},
						},
					},
					{
						Name:        "remove",
						Description: "Removes a role that controls settings",
						Type:        dg.ApplicationCommandOptionSubCommand,
						Options: []*dg.ApplicationCommandOption{
							{
								Name:        "role",
								Description: "The role to remove",
								Required:    true,
								Type:        dg.ApplicationCommandOptionRole,
							},
						},
					},
				},
			},
			{
				Name:        "killmail",
				Description: "The kill mail module settings",
				Type:        dg.ApplicationCommandOptionSubCommandGroup,
				Options: []*dg.ApplicationCommandOption{
					{
						Name:        "channel",
						Description: "Specify which channel to look for kill mails",
						Type:        dg.ApplicationCommandOptionSubCommand,
						Options: []*dg.ApplicationCommandOption{
							{
								Name:        "channel-name",
								Description: "The channel to use",
								Required:    true,
								Type:        dg.ApplicationCommandOptionChannel,
							},
						},
					},
				},
			},
		},
	},
}
