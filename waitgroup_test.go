package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// WaitGroup
// these method is for waiting asynchronous processes to be done
// we dont need to use time.Sleep() for waiting goroutine func to be done
// if using time.Sleep() we actually guessing how many second the goroutines proccess to be done
// so by implementing this WaitGroup, it will wait until all goroutine proccess done before end the func proccess

func RunAsynchronous(group *sync.WaitGroup) {
	// This tell the group to decrement by 1
	defer group.Done()

	// This tell the group to increment by 1
	group.Add(1)

	// add some proccess here
	fmt.Println("This is proccess from func RunAsynchronous")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {

	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	// This will block or wait the all goroutines to finish before finish the func process
	group.Wait()
	fmt.Println("All goroutines complete!")
	
}