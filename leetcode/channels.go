package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	errChan := make(chan error, 1)
	var msg error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Wait()
		select {
		case <-errChan:
			fmt.Printf("\nError read from channel: %v", msg)
		case <-time.After(time.Second * 2):
			fmt.Println("\nTime out: No news in one minute")
		default:
			fmt.Println("Not read from channel")
		}

	}()
	msg = fmt.Errorf("My Error")
	errChan <- msg
	wg.Done()
	time.Sleep(time.Second * 5)
	fmt.Println("\nDone with everything")

}
