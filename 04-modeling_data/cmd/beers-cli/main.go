package main

import (
	"github.com/MGurtD/golang-examples/04-modeling_data/internal/cli"
	"github.com/MGurtD/golang-examples/04-modeling_data/internal/storage/csv"
	"github.com/spf13/cobra"
)

func main() {

	csvRepo := csv.NewRepository()

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(csvRepo))
	rootCmd.Execute()
}
