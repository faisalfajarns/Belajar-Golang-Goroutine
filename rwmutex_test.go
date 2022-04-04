package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.balance = account.balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.balance
	account.RWMutex.RUnlock()

	return balance
}

func TestRWMutex(t *testing.T) {

	account := BankAccount{}
	for i := 0 ; i < 10 ; i++ {
		go func ()  {
			for j := 0 ; j < 10 ; j++ {
				account.AddBalance(1)
				fmt.Println("GET BALANCE : " ,account.GetBalance())
			}	
		}()

	}
	time.Sleep(time.Second * 5)
	fmt.Println("Total : ", account.GetBalance())
}