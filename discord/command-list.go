package discord

import "github.com/bwmarrin/discordgo"

var commandList = []*discordgo.ApplicationCommand{
	{
		Name:        "add-feed",
		Description: "adds rss/atom feed to list",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "link",
				Description: "Link to rss/atom feed",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "channel",
				Description: "Select channel to post new items from new feed",
				Required:    false,
			},
		},
	},
}
