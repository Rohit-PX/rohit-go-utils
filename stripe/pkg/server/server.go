package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rohit-go-utils/stripe/pkg/database"
	"rohit-go-utils/stripe/pkg/errorstripe"
	"sort"
	"syscall"
	"time"
)

func GetTransaction(fileName string) (database.Transactions, error) {
	file, err := os.Open(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, &errorstripe.ErrNotFound{
				ID:   fileName,
				Type: "file",
			}
		}
		return nil, err
	}

	info, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}
	if info.Size() == 0 {
		return nil, &errorstripe.ErrIsEmpty{
			ID:   fileName,
			Type: "file",
		}
	}

	var transactions database.Transactions
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}
func GetTransactionCustomer(fileName, customerID string) (database.CustomerOutput, error) {
	file, err := os.Open(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, &errorstripe.ErrNotFound{
				ID:   fileName,
				Type: "file",
			}
		}
		return nil, err
	}

	info, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}
	if info.Size() == 0 {
		return nil, &errorstripe.ErrIsEmpty{
			ID:   fileName,
			Type: "file",
		}
	}

	var transactions database.Transactions

	defer file.Close()

	if err := json.NewDecoder(file).Decode(&transactions); err != nil {
		return nil, err
	}

	fmt.Println("All transactions", transactions)
	dollarOut := make(map[string]int64)
	for _, tx := range transactions {
		if tx.CustomerId == customerID && tx.Status == "success" {
			dollarOut[customerID] += tx.Amount
		}
	}
	return convertCentsToDollars(dollarOut), nil
}

func GetTotalsPerCurrency(transactions database.Transactions) database.Output {
	totals := make(map[string]int64)
	failures := make(map[string]int)

	for _, tx := range transactions {
		if tx.CustomerId == "" || tx.Currency == "" || tx.Status == "" {
			fmt.Println("skipping this record", tx.CustomerId, tx.Currency, tx.Status)
		} else {
			if tx.Status == "success" {
				totals[tx.Currency] += tx.Amount
			} else {
				failures[tx.CustomerId] += 1

			}
		}
	}

	totalDollars := convertCentsToDollars(totals)
	frequentFailures := getFreqFailures(failures)
	var Output database.Output

	Output.Totals = totalDollars
	Output.FrequentFailures = frequentFailures

	return Output
}
func transactionsHandler(w http.ResponseWriter, r *http.Request) {
	allTransactions, err := GetTransaction("input.json")
	if err != nil {
		log.Panic(err)
	}

	result := GetTotalsPerCurrency(allTransactions)

	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		http.Error(w, fmt.Sprintf("error marshalling output: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func customerTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	cId := r.URL.Query().Get("custid")
	if cId == "" {
		http.Error(w, "empty customer id passed", http.StatusInternalServerError)
	}
	custTransactions, err := GetTransactionCustomer("input.json", cId)
	if err != nil {
		log.Panic(err)
	}

	out, err := json.MarshalIndent(custTransactions, "", "  ")
	if err != nil {
		http.Error(w, fmt.Sprintf("error marshalling output for customers: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func convertCentsToDollars(totalCents map[string]int64) map[string]float64 {
	totalDollars := make(map[string]float64)

	for curr, cents := range totalCents {
		totalDollars[curr] = float64(cents) / 100.0
	}

	return totalDollars
}

func getFreqFailures(failures map[string]int) []database.FailureInfo {
	var freqFailures []database.FailureInfo
	for cid, fails := range failures {
		if fails > 1 {
			freqFailures = append(freqFailures, database.FailureInfo{
				CustomerID: cid,
				Failures:   fails,
			})
		}
	}

	sort.Slice(freqFailures, func(i, j int) bool {
		return freqFailures[i].Failures > freqFailures[j].Failures
	})
	return freqFailures

}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	resp := map[string]string{
		"error":   "Not Found",
		"message": fmt.Sprintf("Route %s %s not found", r.Method, r.URL.Path),
	}

	json.NewEncoder(w).Encode(resp)
}

func Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			notFoundHandler(w, r)
			return
		}
		fmt.Fprintln(w, "Transaction server!")
	})

	mux.HandleFunc("/transactions", transactionsHandler)
	mux.HandleFunc("/customers", customerTransactionsHandler)

	srv := &http.Server{
		Addr:    ":8082",
		Handler: mux,
	}

	go func() {
		fmt.Println("http server started successfully")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error", err)
		}
	}()

	// Listen for interrupt or terminate signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop // Wait until signal is received
	fmt.Println("\nShutting down server...")

	// Gracefully shut down with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown:", err)
	}
	defer cancel()
	return nil

}
