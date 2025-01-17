package main

import (
	"flag"

	"github.com/MGurtD/golang-examples/09-benchmarking/internal/fetching"

	"github.com/MGurtD/golang-examples/09-benchmarking/internal/storage/ontario"

	beerscli "github.com/MGurtD/golang-examples/09-benchmarking/internal"
	"github.com/MGurtD/golang-examples/09-benchmarking/internal/cli"
	"github.com/MGurtD/golang-examples/09-benchmarking/internal/storage/csv"
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

	fetchingService := fetching.NewService(repo)

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(fetchingService))
	rootCmd.Execute()
}
