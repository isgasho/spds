package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var inputFile string
var algorithm string
var size int

var rootCmd = &cobra.Command{
	Use:   "spds",
	Short: "spds is a stream summarizer",
	Long: `Build on top of probabilistic data structures with a small memory footprint spds 
	allows you to process streams with billions of elements in a efficient way.
	
	Complete documentation is available at www.github.com/jomsdev/spds`,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputFile, "file", "f", "", "Input file (required)")
	rootCmd.MarkFlagRequired("file")
	rootCmd.PersistentFlags().StringVarP(&algorithm, "algorithm", "a", "", "Algorithm (Default: KMV)")
	rootCmd.PersistentFlags().IntVarP(&size, "size", "s", 256, "Algorithm")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
