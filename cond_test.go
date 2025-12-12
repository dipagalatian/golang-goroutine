package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker =  sync.Mutex{}
var cond =  sync.NewCond(&locker)
var wg = sync.WaitGroup{}

func WaitCondition (v int) {

	defer wg.Done()
	
	cond.L.Lock()
	cond.Wait()

	fmt.Println("Done", v)

	cond.L.Unlock()
	
}

func TestCond(t *testing.T) {

	// create 100 goroutines
	for i := 0; i < 10; i++ {

		wg.Add(1)
		go WaitCondition(i)

	}

	// Signal()
	// to signal the cond one by one manually for each goroutines
	// go func ()  {
	// 	for i := 0; i < 10; i++ {
			
	// 		fmt.Println("Send signal to cond")
	// 		time.Sleep(1 * time.Second)

	// 		cond.Signal()
		
	// 	}
	// }()

	// Broadcast()
	// to signal the cond one by one manually for each goroutines
	go func ()  {
		fmt.Println("Send broadcast to all goroutines")
		time.Sleep(2 * time.Second)
		cond.Broadcast()
		
	}()

	wg.Wait()
	fmt.Println("All goroutines finish")

}
