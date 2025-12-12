package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// NewTimer
// Manually set the timer and get the object data
// then use the timer channel to get the data
func TestTimer(t *testing.T) {
	
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Time now:", time.Now())

	data := <- timer.C
	fmt.Println("Timer:", data)

}

// After
// More simple then NewTimer
// we can use the returned value which is channel
// just use the channel to get the data
func TestAfter(t *testing.T) {
	
	channel := time.After(5 * time.Second)
	fmt.Println("Time now:", time.Now())

	data := <- channel
	fmt.Println("Timer:", data)
}

// AfterFunc
// we can use this without manage the data or even the channel
// we can pass the func to run after the delay
func TestAfterFunc(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(1)

	// This run as goroutine
	time.AfterFunc(5 * time.Second, func ()  {
		defer wg.Done()

		fmt.Println("Timer:", time.Now())
		
	})

	fmt.Println("Now:", time.Now())

	wg.Wait()
	fmt.Println("Goroutine finish")
	
}