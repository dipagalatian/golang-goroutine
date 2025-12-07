package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// Race condition
// common problem if we dealing with Concurrency or Paraller programming
// use case: 1000 goroutines try to access the same variable called totalCounter
// race condtion happen: multiple goroutines try to increment totalCounter + 1 at the same exactly time
// result: the totalCounter not exactly 1000 at the end of the process
// solution: can see mutex_test.go file
func TestRaceCondition(t *testing.T) {

	totalCounter := 0

	// create 1000 goroutines
	for i := 1; i <=1000; i++ {

		go func ()  {
			// each goroutine will add 100 number to totalNumber
			for j := 1; j <= 100; j++ {
				totalCounter += 1
			}
		}()
		
	}

	time.Sleep(5 * time.Second)
	// expected result totalNumber is 100_000 (if no race condition)
	fmt.Printf("Total counter: %d\n", totalCounter)
	
}