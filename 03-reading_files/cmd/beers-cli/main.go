package main

import (
	"github.com/MGurtD/golang-examples/03-reading_files/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.AddCommand(cli.InitStoreCmd())
	rootCmd.Execute()
}
