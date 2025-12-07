package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Mutex
// Lock & Unlock solution for race condition
// can see the race condition problem on race_condition_test.go file
func TestMutex(t * testing.T) {

	totalCounter := 0
	var mutex sync.Mutex

	// create 1000 goroutines
	for i := 1; i <= 1000; i++ {

		go func ()  {
			// each goroutine will add 100 number to totalNumber
			for j := 1; j <= 100; j++ {

				mutex.Lock() // lock the variable before the adding value
				totalCounter = totalCounter + 1
				mutex.Unlock() //unlock the mutex locking, another goroutines can lock again after this finish
			}
		}()
		
	}

	time.Sleep(5 * time.Second)
	// expected result totalNumber is 100_000 (if no race condition)
	fmt.Printf("Total counter: %d\n", totalCounter)
	
}