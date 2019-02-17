package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func getScanner() *bufio.Scanner {
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	return scanner
}
