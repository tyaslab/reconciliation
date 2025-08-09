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

func getLongestCommonSubsequence(listA, listB []Transaction) *LCSResult {
	lcsResultList := []LCSResult{}

	a := 0
	b := 0

	matchA := -1
	matchB := -1

	count := 0

	for a < len(listA) && b < len(listB) {
		if listA[a].Amount == listB[b].Amount {
			count += 1
			if matchA < 0 {
				matchA = a
			}

			if matchB < 0 {
				matchB = b
			}

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
			if b < (len(listB) - 1) {
				b += 1
			} else {
				a += 1
				b = matchB + 1
			}
			// also reset matchA and B
			matchA = -1
			matchB = -1
		}
	}

	if count > 0 {
		// proceed last matched after iteration ends
		lcsResultList = append(lcsResultList, LCSResult{
			IndexA: matchA,
			IndexB: matchB,
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
