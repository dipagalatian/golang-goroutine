package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// NewTicker
// return object of ticker
// we should get the channel and data manually

func TestTickerStop(t *testing.T) {

	// Create a ticker that ticks every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// Create a channel to signal the worker to stop
	done := make(chan bool)

	// Start the background worker to listen for ticker events
	go func ()  {
		for {
			select {
			case <- done:
				fmt.Println("Worker received stop signal")
				return
			case data := <- ticker.C:
				fmt.Println("Worker received at:", data.Format("15:04:05"))
			}
		}
	}()


	// Let the func wait for 10 seconds before sending the stop signal
	time.Sleep(10 * time.Second)

	// Send the stop signal to the worker
	done <- true
	fmt.Println("Ticker stop")
}

// Tick
// not return the object instead return the channel
func TestTickOnly(t *testing.T) {

	channel := time.Tick(2 * time.Second)

	for data := range channel {
		fmt.Println("Now:", data)
	}
	
}