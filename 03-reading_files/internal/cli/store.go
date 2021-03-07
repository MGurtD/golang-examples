package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// StoreFn function definion of run cobra command
type StoreFn func(cmd *cobra.Command, args []string)

const flag = "id"

// InitStoreCmd initialize beers command
func InitStoreCmd() *cobra.Command {
	storesCmd := &cobra.Command{
		Use:   "store",
		Short: "Print data about store",
		Run:   runStoreFn(),
	}

	storesCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return storesCmd
}

func runStoreFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {

		f, _ := os.Open("03-reading_files/data/stores.csv")
		reader := bufio.NewReader(f)

		var beers = make(map[int]string)

		for line := readLine(reader); line != nil; line = readLine(reader) {
			values := strings.Split(string(line), ",")

			productID, _ := strconv.Atoi(values[0])
			beers[productID] = values[1]
		}

		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			i, _ := strconv.Atoi(id)
			fmt.Println(beers[i])
		} else {
			fmt.Println(beers)
		}
	}
}
