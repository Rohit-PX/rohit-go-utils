package main

import (
	"fmt"
	"math/rand"
	"time"
)

func num(alertChan chan bool) {
	ticker := time.NewTicker(1 * time.Second)

	for t := range ticker.C {
		i := rand.Intn(6)
		fmt.Printf("\nAt T %v I: %d", t, i)
		if i == 5 {
			ticker.Stop()
			fmt.Printf("\nPoison pill.")
			alertChan <- true
		}
	}
}

func program(errChan chan bool) {
	go num(errChan)
	fmt.Printf("\nDoing something.")
	time.Sleep(time.Second * 5)
	fmt.Printf("\nSomething done.")
}

func main() {
	var doneChan chan bool
	var quit chan bool

	program(quit)

	go func() {
		for {
			select {
			//case <-doneChan:
			//	fmt.Println("Alert Done")
			//	return

			case <-quit:
				//ticker.Stop()
				fmt.Println("Quit received.")
				return
			}
		}
	}()

}
