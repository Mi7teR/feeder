package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "feeder",
	Short: "RSS feeds reader bot for Discord",
	Long:  `RSS feeds reader bot for Discord`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
