package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Mutex
// Lock & Unlock solution for race condition
// can see the race condition problem on race_condition_test.go file
func TestMutex(t * testing.T) {

	totalCounter := 0
	var mutex sync.Mutex

	// create 1000 goroutines
	for i := 1; i <= 1000; i++ {

		go func ()  {
			// each goroutine will add 100 number to totalNumber
			for j := 1; j <= 100; j++ {

				mutex.Lock() // lock the variable before the adding value
				totalCounter = totalCounter + 1
				mutex.Unlock() //unlock the mutex locking, another goroutines can lock again after this finish
			}
		}()
		
	}

	time.Sleep(5 * time.Second)
	// expected result totalNumber is 100_000 (if no race condition)
	fmt.Printf("Total counter: %d\n", totalCounter)
	
}

// RW Mutex (Read Write Mutex)
// we can separate locking for operation read and write
// lock & unlock -> write conditioin
// RLock & RUnlock -> read condition
// example: we have bank account struct, we want to add balance and get balance

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (bk *BankAccount) AddBalance(ammount int) {
	bk.RWMutex.Lock()
	bk.Balance = bk.Balance + ammount
	bk.RWMutex.Unlock()
}
func (bk *BankAccount) GetBalance() int {
	bk.RWMutex.RLock()
	balance := bk.Balance
	bk.RWMutex.RUnlock()

	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func ()  {

			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println("Balance now:", account.GetBalance())
			}
			
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total balance:", account.GetBalance())
	
}