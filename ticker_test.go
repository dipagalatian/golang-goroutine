package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// NewTicker
// return object of ticker
// we should get the channel and data manually
func TestTicker(t *testing.T) {

	ticker := time.NewTicker(2 * time.Second)

	go func ()  {
		time.Sleep(10 * time.Second)
		ticker.Stop()
	}()

	for data := range ticker.C {
		fmt.Println("Now:", data)

	}
}

// Tick
// not return the object instead return the channel
func TestTickOnly(t *testing.T) {

	channel := time.Tick(2 * time.Second)

	for data := range channel {
		fmt.Println("Now:", data)
	}
	
}