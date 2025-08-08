package app

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"time"
)

func ReadCSV(stream io.Reader, accountType AccountType, dateFrom, dateTo time.Time) ([]Transaction, error) {
	reader := csv.NewReader(stream)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV file: %v", err)
	}

	data := []Transaction{}
	for _, record := range records {
		var transaction *Transaction
		if accountType == TransactionAccount {
			transaction, err = NewTransaction(record)
			if err != nil {
				return nil, err
			}
		} else {
			transaction, err = NewTransactionFromBankStatement(record)
			if err != nil {
				return nil, err
			}
		}

		if (transaction.TransactionTime.After(dateFrom) || transaction.TransactionTime.Equal(dateFrom)) && (transaction.TransactionTime.Before(dateTo) || transaction.TransactionTime.Equal(dateTo)) {
			data = append(data, *transaction)
		}
	}

	return data, nil
}

func CompareList(transactionAccountList, bankAccountList []Transaction) *Result {
	totalProcessed := len(transactionAccountList) + len(bankAccountList)
	totalMatched := 0
	for {
		lcs := getLongestCommonSubsequence(transactionAccountList, bankAccountList)
		if lcs == nil {
			break
		}

		totalMatched += lcs.Count

		// split transactionAccountList

		transactionAccountList = append(transactionAccountList[0:lcs.IndexA], transactionAccountList[lcs.IndexA+lcs.Count:]...)
		bankAccountList = append(bankAccountList[0:lcs.IndexB], bankAccountList[lcs.IndexB+lcs.Count:]...)
	}

	// total absolute transaction amount
	var transactionDiscrepancy float64 = 0
	for _, t := range transactionAccountList {
		transactionDiscrepancy += math.Abs(t.Amount)
	}

	var bankDiscrepancy float64 = 0
	for _, b := range bankAccountList {
		bankDiscrepancy += math.Abs(b.Amount)
	}

	return &Result{
		Processed:       totalProcessed,
		Matched:         totalMatched,
		SystemUnmatched: len(transactionAccountList),
		BankUnmatched:   len(bankAccountList),
		Discrepancies:   transactionDiscrepancy + bankDiscrepancy,
	}
}

func getLongestCommonSubsequence(transactionListA, transactionListB []Transaction) *LCSResult {
	lcsResultList := []LCSResult{}

	a := 0
	b := 0

	matchA := 0
	matchB := 0

	count := 0

	for a < len(transactionListA) && b < len(transactionListB) {
		if transactionListA[a].Amount == transactionListB[b].Amount {
			count += 1

			// save last matched position
			matchA = a
			matchB = b

			// move a and b when amount equals
			a += 1
			b += 1
		} else {
			if count > 0 {
				lcsResultList = append(lcsResultList, LCSResult{
					IndexA: matchA,
					IndexB: matchB,
					Count:  count,
				})

			}

			count = 0 // reset count when a and b are different
			if b < (len(transactionListB) - 1) {
				b += 1
			} else {
				a += 1
				b = matchB + 1
			}
		}
	}

	if count > 0 {
		// this means all are matched
		lcsResultList = append(lcsResultList, LCSResult{
			IndexA: matchA - 1,
			IndexB: matchB - 1,
			Count:  count,
		})

	}

	if len(lcsResultList) == 0 {
		return nil
	}

	lcsResultIndex := 0
	lastCount := lcsResultList[0].Count

	for i, lcsResult := range lcsResultList {
		if lcsResult.Count > lastCount {
			lcsResultIndex = i
		}
	}

	result := lcsResultList[lcsResultIndex]

	return &result
}
