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
	num(errChan)
	fmt.Printf("\nDoing something.")
	time.Sleep(time.Second * 5)
	fmt.Printf("\nSomething done.")
}

func main() {
	quit := make(chan bool, 1)
	//expired := make(chan bool, 1)

	ticker := time.NewTicker(1 * time.Second)
	startTime := time.Now()

	var x bool
	var y bool
	go func() {
		for _ = range ticker.C {
			i := rand.Intn(6)
			fmt.Printf("\nAt T %v I: %d\n", time.Since(startTime), i)
			if i == 5 {
				x = true
				quit <- true
				fmt.Printf("\nPoison pill.")
			}
		}
	}()

	go func() {
		time.Sleep(time.Second * 10)
		quit <- true
		y = true
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Ticking")

		case <-quit:
			fmt.Println("Quit received.")
		}
		if x {
			fmt.Println("Number found.")
			break
		} else if y {
			fmt.Println("Timer expired")
			break
		}
	}

}
