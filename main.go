package main

import (
	"github.com/mittacy/ego-layout/cmd/start"
	"github.com/spf13/cobra"
	"log"
)

const version = "v1.0.0"

var rootCmd = &cobra.Command{
	Use:     "server",
	Short:   "server: An elegant toolkit for start server.",
	Long:    "server: An elegant toolkit for start server.",
	Version: version,
}

func init() {
	rootCmd.AddCommand(start.Cmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
