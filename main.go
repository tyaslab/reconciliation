package main

import (
	"fmt"
	"os"
	"time"

	"adityadarmawan.id/recon/app"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Printf("Usage: %s <fileA.csv> <fileB.csv> <dateFrom> <dateTo>\n", os.Args[0])
		os.Exit(1)
	}

	fileA := os.Args[1]
	fileB := os.Args[2]

	dateFrom, err := time.Parse(app.DATE_TIME_LAYOUT, os.Args[3])
	if err != nil {
		fmt.Printf("Cannot parse date from: %s\n", os.Args[3])
		os.Exit(1)
	}

	dateTo, err := time.Parse(app.DATE_TIME_LAYOUT, os.Args[4])
	if err != nil {
		fmt.Printf("Cannot parse date to: %s\n", os.Args[4])
		os.Exit(1)
	}

	dataA, err := app.ReadCSV(fileA, app.TransactionAccount, dateFrom, dateTo)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", fileA, err)
		os.Exit(1)
	}

	dataB, err := app.ReadCSV(fileB, app.BankAccount, dateFrom, dateTo)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", fileB, err)
		os.Exit(1)
	}

	result := app.CompareList(dataA, dataB)

	printResult(result)
}

func printResult(result *app.Result) {
	fmt.Printf("total processed: %d\n", result.Processed)
	fmt.Printf("total matched: %d\n", result.Matched)
	fmt.Printf("total system transaction unmatched: %d\n", result.SystemUnmatched)
	fmt.Printf("total bank unmatched: %d\n", result.BankUnmatched)
	fmt.Printf("total discrepancies: %f\n", result.Discrepancies)
}
