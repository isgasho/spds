package cmd

import (
	"fmt"
	"spds/pkg/sample"

	"github.com/spf13/cobra"
)

var sampleCmd = &cobra.Command{
	Use:     "sample",
	Short:   "Returns a sample of a stream",
	Example: "spds sample -s {size} -f {path/to/file}",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := getScanner()
		sampler := sample.NewAffirmative(size)

		for scanner.Scan() {
			sampler.Add(scanner.Text())
		}

		msg := "Sample size of %d out of %d elements (counting repetitions)"
		template := fmt.Sprintf(msg, sampler.Size(), sampler.ElementsAdded())
		fmt.Println(template)

		for _, value := range sampler.Sample() {
			fmt.Println(value)
		}
	},
}

func init() {
	rootCmd.AddCommand(sampleCmd)
}
