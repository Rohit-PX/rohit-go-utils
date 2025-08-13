package main

import (
	"fmt"
	"log"
	"rohit-go-utils/stripe/cmd/server"
)

func main() {
	allTransactions, err := server.GetTransaction("input.json")
	if err != nil {
		log.Panic(err)
	}

	totals := server.GetTotalsPerCurrency(allTransactions)

	fmt.Println("Total per currency", totals)
}
