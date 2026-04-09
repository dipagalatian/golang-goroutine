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

// sync.Once is a struct that will make sure the function only execute once, even if we call it multiple times
// this is useful for example when we want to initialize a resource that only need to be initialized once, like database connection, or configuration file
// if we call the function multiple times, it will only execute the first time, and ignore the rest of the calls
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

type PromotionCode struct {
	mu sync.Mutex
	IsUsed bool
	Winner string
}

func (p *PromotionCode) Use(username string) (bool, string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.IsUsed {
		// fmt.Println("Promotion code already used by", p.Winner)
		return false, p.Winner
	}
	
	p.IsUsed = true
	p.Winner = username
	// fmt.Printf("Congratulations %s! You have used the promotion code successfully.\n", username)
	return true, username
	
}

func TestBookWithPromotionCode(t *testing.T) {
	promotionCode := &PromotionCode{}

	var wg sync.WaitGroup

	// Simulate multiple users trying to use the promotion code concurrently
	for i := 0; i < 30; i++ {
		wg.Add(1)

		username := fmt.Sprintf("User%d", i+1)

		go func (name string)  {
			defer wg.Done()
			used, _ := promotionCode.Use(name)

			if used {
				fmt.Printf(">>> %s successfully used the promotion code!\n", name)
			} 

		}(username)
	}

	wg.Wait()

	fmt.Printf("Promotion code state: Used=%v, Winner=%s\n", promotionCode.IsUsed, promotionCode.Winner)

	if !promotionCode.IsUsed {
		t.Error("Promotion code should be used but it is not")
	}
	if promotionCode.Winner == "" {
		t.Error("Promotion code should have a winner but it does not")
	}
}