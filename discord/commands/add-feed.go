package commands

import (
	"fmt"
	"log"
	"net/url"

	"github.com/bwmarrin/discordgo"
)

func AddFeed(s *discordgo.Session, i *discordgo.InteractionCreate) {
	feedSourceLink := i.ApplicationCommandData().Options[0].StringValue()
	feedTargetChannel := i.ApplicationCommandData().Options[1].ChannelValue(s)
	if feedTargetChannel != nil {
		s.ChannelMessageSend(feedTargetChannel.ID, fmt.Sprintf("Adding feed %s", feedSourceLink))
	}

	response := new(discordgo.InteractionResponse)
	response.Type = discordgo.InteractionResponseChannelMessageWithSource

	if !isURL(feedSourceLink) {
		response.Data = &discordgo.InteractionResponseData{
			Content: "Invalid URL",
		}
	} else {
		response.Data = &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Added feed: %s", feedSourceLink),
		}
	}
	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		log.Println(err)
	}
}

// isURL - check string is valid url
func isURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}
