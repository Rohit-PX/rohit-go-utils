package server

import (
	"encoding/json"
	"os"
	"rohit-go-utils/stripe/cmd/database"
)

func GetTransaction(fileName string) (database.Transactions, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	var transactions database.Transactions
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func GetTotalsPerCurrency(transactions database.Transactions) map[string]int64 {
	totals := make(map[string]int64)

	for _, tx := range transactions {
		if _, ok := totals[tx.Currency]; !ok {
			totals[tx.Currency] = tx.Amount
		} else {
			totals[tx.Currency] += tx.Amount
		}
	}
	return totals
}

func GetFailures(transactions database.Transactions) database.FailureInfo {
	return database.FailureInfo{}
}
