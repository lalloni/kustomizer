package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var VERSION = "0.0.0-dev.0"

var rootCmd = &cobra.Command{
	Use:           "kustomizer",
	Version:       VERSION,
	SilenceUsage:  true,
	SilenceErrors: true,
	Short:         "A command line utility for generating, building and applying kustomizations",
}

var (
	kubectl string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&kubectl, "kubectl", "kubectl", "Command to run kubectl")
}

func main() {
	log.SetFlags(0)
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
