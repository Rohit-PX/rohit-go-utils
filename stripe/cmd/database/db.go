package database

type Transaction struct {
	CustomerId string `json:"customer_id,omitempty"`
	Amount     int64  `json:"amount,omitempty"`
	Currency   string `json:"currency,omitempty"`
	Status     string `json:"status,omitempty"`
}

type Transactions []Transaction

type FailureInfo struct {
	CustomerID string `json:"customer_id"`
	Failures   int    `json:"failures"`
}

type Output struct {
	Totals           map[string]float64 `json:"totals"`
	FrequentFailures []FailureInfo      `json:"frequent_failures"`
}
