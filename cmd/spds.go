package main

import (
	"bufio"
	"fmt"
	"os"
	"spds/pkg/cardinality"
)

func main() {
	d := cardinality.NewDummy()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		d.Add(scanner.Text())
	}

	fmt.Println("size: ", int(d.EstimateCardinality()))
}
