package app

import (
	"encoding/json"
	"net/http"
	"time"
)

func HandlerReconcile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// filter date
	dateFrom := r.FormValue("date_from")
	if dateFrom == "" {
		http.Error(w, "date_from is required", http.StatusBadRequest)
		return
	}

	dateTo := r.FormValue("date_to")
	if dateTo == "" {
		http.Error(w, "date_to is required", http.StatusBadRequest)
		return
	}

	dtFrom, err := time.Parse(DATE_LAYOUT, dateFrom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dtTo, err := time.Parse(DATE_LAYOUT, dateTo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dtTo = time.Date(dtTo.Year(), dtTo.Month(), dtTo.Day(), 23, 59, 59, 999_999_999, dtTo.Location())

	// transaction
	transactionFile, _, err := r.FormFile("transaction_file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer transactionFile.Close()

	transaction, err := ReadCSV(transactionFile, TransactionAccount, dtFrom, dtTo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// bank
	bankStatementFile, _, err := r.FormFile("bank_statement_file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer bankStatementFile.Close()

	bankStatement, err := ReadCSV(bankStatementFile, BankAccount, dtFrom, dtTo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// compare
	result := CompareList(transaction, bankStatement)

	resultJson, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
}
