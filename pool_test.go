package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {

	connectionPool := sync.Pool{
		New: func() any {
			return "New"
		},
	}
	wg := &sync.WaitGroup{}

	connectionPool.Put("dipa")
	connectionPool.Put("galatian")
	connectionPool.Put("parar")

	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func (i int)  {
			defer wg.Done()

			data := connectionPool.Get()
			fmt.Println("data connectionPool:", data, "index:", i)
			time.Sleep(2 * time.Second)
			connectionPool.Put(data)
			
		}(i)

	}

	wg.Wait()
	fmt.Println("TestPoll finish")
	
}
