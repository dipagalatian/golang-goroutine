package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {

	once := sync.Once{}
	group := sync.WaitGroup{}

	// 100 goroutine will try to call the OnlyOnce func
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func ()  {
			defer group.Done()
			once.Do(OnlyOnce)
			
		}()
	}

	group.Wait()
	fmt.Println("func execution done, total counter:", counter)
	
}