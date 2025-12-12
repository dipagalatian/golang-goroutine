package golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {

	var counter int64 = 0
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func ()  {
			defer wg.Done()
			
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
			
		}()
	}

	wg.Wait()
	fmt.Println("All goroutines finish with total counter:", counter)
	
	
}