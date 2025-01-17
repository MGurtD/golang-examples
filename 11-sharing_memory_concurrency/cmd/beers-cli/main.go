package main

import (
	"flag"

	beerscli "github.com/MGurtD/golang-examples/11-sharing_memory_concurrency/internal"
	"github.com/MGurtD/golang-examples/11-sharing_memory_concurrency/internal/cli"
	"github.com/MGurtD/golang-examples/11-sharing_memory_concurrency/internal/fetching"
	"github.com/MGurtD/golang-examples/11-sharing_memory_concurrency/internal/storage/csv"
	"github.com/MGurtD/golang-examples/11-sharing_memory_concurrency/internal/storage/ontario"
	"github.com/spf13/cobra"
)

func main() {
	// CPU profiling code starts here
	//f, _ := os.Create("beers.cpu.prof")
	//defer f.Close()
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()
	// CPU profiling code ends here

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

	// Memory profiling code starts here
	//f2, _ := os.Create("beers.mem.prof")
	//defer f2.Close()
	//pprof.WriteHeapProfile(f2)
	// Memory profiling code ends here
}
