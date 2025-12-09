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

func (ba *BankAccount) AddBalance(ammount int) {
	ba.RWMutex.Lock()
	ba.Balance = ba.Balance + ammount
	ba.RWMutex.Unlock()
}
func (ba *BankAccount) GetBalance() int {
	ba.RWMutex.RLock()
	balance := ba.Balance
	ba.RWMutex.RUnlock()

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

// Deadlock
// deadlock is happen when goroutines are waiting for another goroutines to finish, but the other goroutines is waiting for this goroutines to finish
// example: goroutine A is waiting for goroutine B to finish, but goroutine B is waiting for goroutine A to finish
// to solve this problem, we can use channel to communicate between goroutines

type UserBalance struct {
	sync.Mutex
	Name string
	Balance int
}

func (ub *UserBalance) Lock() {
	ub.Mutex.Lock()
}

func (ub *UserBalance) Unlock() {
	ub.Mutex.Unlock()
}

func (ub *UserBalance) Change(amount int) {
	ub.Balance = ub.Balance + amount
}

func Transfer(userFrom *UserBalance, userTo *UserBalance, amount int) {

	userFrom.Lock()
	fmt.Println("Lock userFrom:", userFrom.Name)
	userFrom.Change(-amount)
	// userFrom.Unlock()

	time.Sleep(1 * time.Second)

	userTo.Lock()
	fmt.Println("Lock userTo:", userTo.Name)
	userTo.Change(amount)
	// userTo.Unlock()

	time.Sleep(1 * time.Second)

	userFrom.Unlock()
	userTo.Unlock()
}

func TestDeadlock(t *testing.T) {

	user1 := UserBalance {
		Name: "Dipa",
		Balance: 1_000_000,
	}
	user2 := UserBalance {
		Name: "Galatian",
		Balance: 1_000_000,
	}

	go Transfer(&user1, &user2, 100_000) // goroutine 1
	go Transfer(&user2, &user1, 200_000) // goroutine 2

	time.Sleep(3 * time.Second)

	fmt.Println("User name:", user1.Name, "Balance:", user1.Balance)
	fmt.Println("User name:", user2.Name, "Balance:", user2.Balance)	

	// These logs will show balance for user1 is 900_000 and user2 is 800_000
	// question is, why the total money gone 300_000 in these 2 transaction?
	// its because deadlock -> goroutine 1 wait for goroutine 2 and vice versa
	// you can see the Lock logs only show for the "userFrom" in both user
	// so the deadlock flow are:
	// goroutine 1 is locking the userFrom (galatian)
	// goroutine 2 is locking the userFrom (dipa)
	// when goroutine 1 is going to lock the userTo (galatian), its already been locked by goroutine 2
	// and at the same time, when goroutine 2 is going to lock the userTo(dipa), its already been locked by goroutine 1
	// both goroutine is waiting each other and make DEADLOCK happen.
	
}