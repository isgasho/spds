package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of spds",
	Long:  `spds follows Semantic's Versioning specification.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Spds version 0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
