package main

import (
	"flag"

	"github.com/MGurtD/golang-examples/06-error_handling/internal/storage/ontario"

	beerscli "github.com/MGurtD/golang-examples/06-error_handling/internal"
	"github.com/MGurtD/golang-examples/06-error_handling/internal/cli"
	"github.com/MGurtD/golang-examples/06-error_handling/internal/storage/csv"
	"github.com/spf13/cobra"
)

func main() {

	csvData := flag.Bool("csv", false, "load data from csv")
	flag.Parse()
	var repo beerscli.BeerRepo
	repo = csv.NewRepository()

	if !*csvData {
		repo = ontario.NewOntarioRepository()
	}

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(repo))
	rootCmd.Execute()
}
