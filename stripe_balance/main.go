package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var rates = map[string]float64{
	"USD": 1.0,
	"EUR": 1.1,
	"INR": 0.012,
}

//type NetBalance map[string]map[string]float64

type NetBalance map[string]Balance

type BalanceData struct {
	Amount   float64
	Customer string
	Currency string
}

type CustomerData []BalanceData

type Balance map[string]float64

func main() {
	input := "t1:c1:USD:100,t2:c2:EUR:200,t3:c1:USD:-50,t4:c3:INR:5000,t5:c1:USD:250"

	nB, custData, err := ParseTransactions(input)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("NetBalance", nB)

	usdBalance := CurrencyConverter(nB, rates)
	fmt.Println("USD Balance", usdBalance)

	maxBalance := MaxBalance(usdBalance)
	fmt.Println("max Balance", fmt.Sprintf("%.2f", maxBalance))

	fmt.Println("Customer Data", custData)

}

func ParseTransactions(s string) (NetBalance, CustomerData, error) {
	allTransactions := strings.Split(s, ",")
	if len(allTransactions) < 1 {
		return nil, nil, fmt.Errorf("empty input")
	}

	netBalance := make(NetBalance)
	var customerData = CustomerData{}
	for _, trans := range allTransactions {
		t := strings.Split(trans, ":")
		if len(t) < 4 {
			return nil, nil, fmt.Errorf("invalid transaction format")
		}
		custId := t[1]
		currency := t[2]
		amount, err := strconv.Atoi(t[3])

		currBalanceData := BalanceData{Amount: float64(amount), Customer: custId, Currency: currency}
		customerData = append(customerData, currBalanceData)

		fmt.Println("current transaction", trans)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid transaction format")
		}

		// populate the netbalance map
		if _, ok := netBalance[custId][currency]; !ok {
			netBalance[custId] = make(map[string]float64)
		}
		netBalance[custId][currency] += float64(amount)
	}

	return netBalance, customerData, nil

}

func CurrencyConverter(nB NetBalance, rates map[string]float64) Balance {
	balance := make(Balance)
	for c, n := range nB {
		for curr, amt := range n {
			balance[c] = amt * rates[curr]
		}
	}
	return balance
}

func MaxBalance(bal Balance) float64 {
	max := float64(0)
	for _, balance := range bal {
		if balance > max {
			max = balance
		}
	}
	return max
}
