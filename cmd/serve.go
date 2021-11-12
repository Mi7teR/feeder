package cmd

import (
	"log"

	"github.com/Mi7teR/feeder/discord"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run RSS feed reader Discord bot",
	Long:  `Run RSS feed reader Discord bot`,
	Run: func(cmd *cobra.Command, args []string) {
		token, err := cmd.Flags().GetString("token")
		if err != nil {
			log.Fatalln(err)
		}
		err = discord.RunServer(token)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().String("token", "", "Discord bot token")
}
