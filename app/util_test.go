package app

import (
	"os"
	"testing"
	"time"
)

var (
	dateFrom string = "2025-08-01 00:00:00"
	dateTo   string = "2025-08-01 23:59:59"
)

func TestReadCSVAllMatched(t *testing.T) {
	transactionFile, err := os.Open("../fixture/matched/transaction.csv")
	if err != nil {
		t.Errorf("cannot open transaction file: %v", err)
	}

	defer transactionFile.Close()

	dateFromDt, _ := time.Parse(DATE_TIME_LAYOUT, dateFrom)
	dateToDt, _ := time.Parse(DATE_TIME_LAYOUT, dateTo)

	transaction, err := ReadCSV(transactionFile, TransactionAccount, dateFromDt, dateToDt)
	if err != nil {
		t.Errorf("cannot read csv transaction: %v", err)
	}

	bankAccountFile, err := os.Open("../fixture/matched/bank.csv")
	if err != nil {
		t.Errorf("cannot open bank account file: %v", err)
	}

	defer bankAccountFile.Close()

	bankAccount, err := ReadCSV(bankAccountFile, BankAccount, dateFromDt, dateToDt)
	if err != nil {
		t.Errorf("cannot read csv bank account: %v", err)
	}

	result := CompareList(transaction, bankAccount)

	if result.Matched != 4 {
		t.Errorf("Matched should be %d not %d", 4, result.Matched)
	}

	if result.BankUnmatched != 0 {
		t.Errorf("Bank unmatched should be %d not %d", 0, result.BankUnmatched)
	}

	if result.SystemUnmatched != 0 {
		t.Errorf("System unmatched should be %d not %d", 0, result.SystemUnmatched)
	}

	if result.Processed != 8 {
		t.Errorf("Processed should be %d not %d", 8, result.Processed)
	}

	if result.Discrepancies != 0 {
		t.Errorf("Discrepancies should be %d not %f", 0, result.Discrepancies)
	}
}

func TestReadCSVUnmatched(t *testing.T) {
	transactionFile, err := os.Open("../fixture/unmatched/transaction.csv")
	if err != nil {
		t.Errorf("cannot open transaction file: %v", err)
	}

	defer transactionFile.Close()

	dateFromDt, _ := time.Parse(DATE_TIME_LAYOUT, dateFrom)
	dateToDt, _ := time.Parse(DATE_TIME_LAYOUT, dateTo)

	transaction, err := ReadCSV(transactionFile, TransactionAccount, dateFromDt, dateToDt)
	if err != nil {
		t.Errorf("cannot read csv transaction: %v", err)
	}

	bankAccountFile, err := os.Open("../fixture/unmatched/bank.csv")
	if err != nil {
		t.Errorf("cannot open bank account file: %v", err)
	}

	defer bankAccountFile.Close()

	bankAccount, err := ReadCSV(bankAccountFile, BankAccount, dateFromDt, dateToDt)
	if err != nil {
		t.Errorf("cannot read csv bank account: %v", err)
	}

	result := CompareList(transaction, bankAccount)

	if result.Matched != 3 {
		t.Errorf("Matched should be %d not %d", 3, result.Matched)
	}

	if result.BankUnmatched != 1 {
		t.Errorf("Bank unmatched should be %d not %d", 1, result.BankUnmatched)
	}

	if result.SystemUnmatched != 1 {
		t.Errorf("System unmatched should be %d not %d", 1, result.SystemUnmatched)
	}

	if result.Processed != 8 {
		t.Errorf("Processed should be %d not %d", 8, result.Processed)
	}

	if result.Discrepancies != 900 {
		t.Errorf("Discrepancies should be %d not %f", 900, result.Discrepancies)
	}
}

func TestReadCSVUnmatchedDiffCount(t *testing.T) {
	transactionFile, err := os.Open("../fixture/unmatched-diff-count/transaction.csv")
	if err != nil {
		t.Errorf("cannot open transaction file: %v", err)
	}

	defer transactionFile.Close()

	dateFromDt, _ := time.Parse(DATE_TIME_LAYOUT, dateFrom)
	dateToDt, _ := time.Parse(DATE_TIME_LAYOUT, dateTo)

	transaction, err := ReadCSV(transactionFile, TransactionAccount, dateFromDt, dateToDt)
	if err != nil {
		t.Errorf("cannot read csv transaction: %v", err)
	}

	bankAccountFile, err := os.Open("../fixture/unmatched-diff-count/bank.csv")
	if err != nil {
		t.Errorf("cannot open bank account file: %v", err)
	}

	defer bankAccountFile.Close()

	bankAccount, err := ReadCSV(bankAccountFile, BankAccount, dateFromDt, dateToDt)
	if err != nil {
		t.Errorf("cannot read csv bank account: %v", err)
	}

	result := CompareList(transaction, bankAccount)

	if result.Matched != 3 {
		t.Errorf("Matched should be %d not %d", 3, result.Matched)
	}

	if result.BankUnmatched != 3 {
		t.Errorf("Bank unmatched should be %d not %d", 3, result.BankUnmatched)
	}

	if result.SystemUnmatched != 1 {
		t.Errorf("System unmatched should be %d not %d", 1, result.SystemUnmatched)
	}

	if result.Processed != 10 {
		t.Errorf("Processed should be %d not %d", 10, result.Processed)
	}

	if result.Discrepancies != 7900 {
		t.Errorf("Discrepancies should be %d not %f", 7900, result.Discrepancies)
	}
}

func TestReadCSVUnmatchedDiffCountLongTrans(t *testing.T) {
	transactionFile, err := os.Open("../fixture/unmatched-diff-count-long-trans/transaction.csv")
	if err != nil {
		t.Errorf("cannot open transaction file: %v", err)
	}

	defer transactionFile.Close()

	dateFromDt, _ := time.Parse(DATE_TIME_LAYOUT, dateFrom)
	dateToDt, _ := time.Parse(DATE_TIME_LAYOUT, dateTo)

	transaction, err := ReadCSV(transactionFile, TransactionAccount, dateFromDt, dateToDt)
	if err != nil {
		t.Errorf("cannot read csv transaction: %v", err)
	}

	bankAccountFile, err := os.Open("../fixture/unmatched-diff-count-long-trans/bank.csv")
	if err != nil {
		t.Errorf("cannot open bank account file: %v", err)
	}

	defer bankAccountFile.Close()

	bankAccount, err := ReadCSV(bankAccountFile, BankAccount, dateFromDt, dateToDt)
	if err != nil {
		t.Errorf("cannot read csv bank account: %v", err)
	}

	result := CompareList(transaction, bankAccount)

	if result.Matched != 3 {
		t.Errorf("Matched should be %d not %d", 3, result.Matched)
	}

	if result.BankUnmatched != 1 {
		t.Errorf("Bank unmatched should be %d not %d", 1, result.BankUnmatched)
	}

	if result.SystemUnmatched != 3 {
		t.Errorf("System unmatched should be %d not %d", 3, result.SystemUnmatched)
	}

	if result.Processed != 10 {
		t.Errorf("Processed should be %d not %d", 10, result.Processed)
	}

	if result.Discrepancies != 7900 {
		t.Errorf("Discrepancies should be %d not %f", 7900, result.Discrepancies)
	}
}

func TestReadCSVAllMatchedFiltered(t *testing.T) {
	transactionFile, err := os.Open("../fixture/matched-filtered/transaction.csv")
	if err != nil {
		t.Errorf("cannot open transaction file: %v", err)
	}

	defer transactionFile.Close()

	dateFromDt, _ := time.Parse(DATE_TIME_LAYOUT, dateFrom)
	dateToDt, _ := time.Parse(DATE_TIME_LAYOUT, dateTo)

	transaction, err := ReadCSV(transactionFile, TransactionAccount, dateFromDt, dateToDt)
	if err != nil {
		t.Errorf("cannot read csv transaction: %v", err)
	}

	bankAccountFile, err := os.Open("../fixture/matched-filtered/bank.csv")
	if err != nil {
		t.Errorf("cannot open bank account file: %v", err)
	}

	defer bankAccountFile.Close()

	bankAccount, err := ReadCSV(bankAccountFile, BankAccount, dateFromDt, dateToDt)
	if err != nil {
		t.Errorf("cannot read csv bank account: %v", err)
	}

	result := CompareList(transaction, bankAccount)

	if result.Matched != 2 {
		t.Errorf("Matched should be %d not %d", 2, result.Matched)
	}

	if result.BankUnmatched != 0 {
		t.Errorf("Bank unmatched should be %d not %d", 0, result.BankUnmatched)
	}

	if result.SystemUnmatched != 0 {
		t.Errorf("System unmatched should be %d not %d", 0, result.SystemUnmatched)
	}

	if result.Processed != 4 {
		t.Errorf("Processed should be %d not %d", 8, result.Processed)
	}

	if result.Discrepancies != 0 {
		t.Errorf("Discrepancies should be %d not %f", 0, result.Discrepancies)
	}
}
