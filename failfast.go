import "sync"
import "time"

func deleteClusterPairsInParallel(cmDests []int) <-chan error {
	var wg sync.WaitGroup
	deleteErrChan := make(chan error)
	out := make(chan error)
	quit := make(chan bool)

	deletePair := func(cm int, wg *sync.WaitGroup, c chan error) {
		defer wg.Done()
		defer close(quit)
		err := cm % 2
		if err != 0 {
			fmt.Printf("\nERROR %d", cm)
			c <- fmt.Errorf("Error in deleting cluster pair with %d. Error: %v", cm, err)
		} else if cm == 20 {
			fmt.Printf("\nLAST %d", cm)

			out <- nil
			quit <- true
		} else {
			fmt.Printf("\nNO ERROR %d", cm)
		}
		for n := range c {
			out <- n
			if n != nil {
				out <- n
				quit <- true
			}
		}
		select {
		case <-quit:
			fmt.Printf("Quit signal received")
			return
		default:
		}
	}

	wg.Add(len(cmDests))
	for _, dest := range cmDests {
		go deletePair(dest, &wg, deleteErrChan)
	}
	go func() {
		wg.Wait()
	}()

	return out
}

func main() {

	a := []int{4, 8, 10, 11, 16, 18, 20}
	start := time.Now()
	err := <-deleteClusterPairsInParallel(a)
	if err != nil {
		fmt.Printf("\nERROR FOUND: %v", err)
	}
	fmt.Printf("\nTotal time: %v", time.Since(start))
}
