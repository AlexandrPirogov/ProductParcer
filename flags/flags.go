package flags

import (
	"log"

	"github.com/spf13/cobra"
)

// file to read
var Filepath string

var parserCommand = &cobra.Command{
	Use:   "gopars",
	Short: "gopars parses given file and print products with best rating and price",
	Long: `gopars parses given file and print products with best rating and price.
	Supported extensions: .csv, .json.
	`,
}

func ReadConfig() {
	parserCommand.PersistentFlags().StringVarP(&Filepath, "filepath", "f", "", "path to thef file to parse")
	if err := parserCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
