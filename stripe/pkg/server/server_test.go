package server_test

import (
	"rohit-go-utils/stripe/pkg/database"
	"rohit-go-utils/stripe/pkg/server"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTransaction(t *testing.T) {
	txs := database.Transactions{
		{CustomerId: "c1", Amount: 2000, Currency: "usd", Status: "success"},
		{CustomerId: "c1", Amount: 3000, Currency: "eur", Status: "success"},
		{CustomerId: "c2", Amount: 1500, Currency: "usd", Status: "failed"},
		{CustomerId: "c2", Amount: 1545, Currency: "usd", Status: "failed"},
	}

	result := server.GetTotalsPerCurrency(txs)

	assert.Equal(t, len(result.Totals), 2)
	assert.Equal(t, len(result.FrequentFailures), 1)
	assert.Equal(t, result.Totals["usd"], 20.00)

	txs = database.Transactions{
		{CustomerId: "c1", Amount: 2000, Currency: "usd", Status: "success"},
		{CustomerId: "c1", Amount: 3000, Currency: "eur", Status: "success"},
		{CustomerId: "c1", Amount: 1500, Currency: "usd", Status: "failed"},
		{CustomerId: "c2", Amount: 1545, Currency: "eur", Status: "failed"},
	}
	result = server.GetTotalsPerCurrency(txs)

	assert.Equal(t, len(result.Totals), 2)
	assert.Equal(t, len(result.FrequentFailures), 0)

	txs = append(txs, database.Transaction{CustomerId: "c1", Amount: 1555, Currency: "usd", Status: "success"})
	result = server.GetTotalsPerCurrency(txs)
	assert.Equal(t, result.Totals["usd"], 35.55)

}
