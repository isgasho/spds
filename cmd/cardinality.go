package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"spds/pkg/cardinality"
)

func formatCardinality(estimator cardinality.CardinalityEstimator) string {
	distinct := estimator.EstimateCardinality()
	total := estimator.ElementsAdded()
	return fmt.Sprintf("%d distinct elements out of %d processed elements", distinct, total)
}

func getCardinalityEstimator() cardinality.CardinalityEstimator {
	switch algorithm {
	case "dummy":
		return cardinality.NewDummy()
	case "kmv":
		return cardinality.NewKMV(size)
	case "hll":
		return cardinality.NewHLL(size)
	default:
		err := fmt.Sprintf("%s is not a valid algorithm", algorithm)
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}

var cardinalityCmd = &cobra.Command{
	Use:     "cardinality",
	Short:   "Estimates the cardinality (number of distinct elements) of a stream",
	Example: "spds carinality -a {dummy, kmv, hll, hllpp, recordinality} -s {size} -f {path/to/file}",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := getScanner()
		estimator := getCardinalityEstimator()

		for scanner.Scan() {
			estimator.Add(scanner.Text())
		}

		fmt.Println(formatCardinality(estimator))
	},
}

func init() {
	rootCmd.AddCommand(cardinalityCmd)
}
