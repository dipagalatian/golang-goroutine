package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello World")
}


// To run func with goroutines, we can add keyword "go" before call a func
// This execution of TestCreateGoroutines will run asynchronous
// If the goroutines still not finish yet, it will not wait for the goroutines to finish
// Instead it will continue execution the next logic
// time.Sleep() used for make sure goroutines have enough time to run before the func stop the execution
func TestCreateGoroutine(t *testing.T) {
	go HelloWorld()
	fmt.Println("Goroutine Test Done")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(n int) {
	fmt.Println("Goroutine:", n)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}