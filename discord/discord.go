package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Mi7teR/feeder/discord/commands"
	"github.com/bwmarrin/discordgo"
)

func RunServer(token string) error {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("error creating Discord session: %w", err)
	}

	oldCommands, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		return fmt.Errorf("error getting commands: %w", err)
	}

	for _, v := range oldCommands {
		err = s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			return fmt.Errorf("error deleting command: %w", err)
		}
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		commands.AddFeed(s, i)
	})

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})

	err = s.Open()
	if err != nil {
		return fmt.Errorf("error opening Discord session: %w", err)
	}

	for _, v := range commandList {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			return fmt.Errorf("error creating command: %w", err)
		}
	}

	defer func(s *discordgo.Session) {
		err := s.Close()
		if err != nil {
			log.Println("error closing Discord session:", err)
		}
	}(s)

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdown")
	return nil
}
