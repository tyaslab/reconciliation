package app

import (
	"strconv"
	"time"
)

type Result struct {
	Processed       int
	Matched         int
	SystemUnmatched int
	BankUnmatched   int
	Discrepancies   float64
}

type EntryType string

const (
	Debit  EntryType = "DEBIT"
	Credit EntryType = "CREDIT"
)

type AccountType string

const (
	TransactionAccount AccountType = "TRANSACTION"
	BankAccount        AccountType = "BANK"
)

type Transaction struct {
	TrxID           string
	Amount          float64
	EntryType       EntryType
	TransactionTime time.Time
}

func NewTransaction(data []string) (*Transaction, error) {
	amount, err := strconv.ParseFloat(data[1], 64)
	if err != nil {
		return nil, err
	}

	entryType := EntryType(data[2])
	if entryType == Debit {
		amount = 0 - amount
	}

	transactionTime, err := time.Parse(DATE_TIME_LAYOUT, data[3])
	if err != nil {
		return nil, err
	}

	transaction := Transaction{
		TrxID:           data[0],
		Amount:          amount,
		EntryType:       entryType,
		TransactionTime: transactionTime,
	}

	return &transaction, nil
}

func NewTransactionFromBankStatement(data []string) (*Transaction, error) {
	amount, err := strconv.ParseFloat(data[1], 64)
	if err != nil {
		return nil, err
	}

	var entryType EntryType
	if amount < 0 {
		entryType = Debit
	} else {
		entryType = Credit
	}

	transactionTime, err := time.Parse(DATE_TIME_LAYOUT, data[2])
	if err != nil {
		return nil, err
	}

	transaction := Transaction{
		TrxID:           data[0],
		Amount:          amount,
		EntryType:       entryType,
		TransactionTime: transactionTime,
	}

	return &transaction, nil
}

type LCSResult struct {
	IndexA int
	IndexB int
	Count  int
}
