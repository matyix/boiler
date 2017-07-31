package command

import (
	"log"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use: "Sample",
	Run: run,
}

func RootCommand() *cobra.Command {
	rootCmd.PersistentFlags().StringP("Config", "c", "", "config file")
	rootCmd.Flags().IntP("Port", "p", 0, "port")

	return &rootCmd
}

func run(command *cobra.Command, args []string) {
	config, err := conf.LoadConfig(command)
	if err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}
	
	logger, err := conf.ConfigureLogging(&config.LogConfig)
	if err != nil {
		log.Fatal("Failed to configure logging: " + err.Error())
	}

	logger.Infof("Starting with config: %+v", config)
}
