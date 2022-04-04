package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Locks(){
	user.Mutex.Lock()
}

func (user *UserBalance) Unlocks(){
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int){
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int){
	user1.Locks()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)
	time.Sleep(time.Second *2)
	
	user2.Locks()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)
	time.Sleep(time.Second *2)
	
	user1.Unlocks()
	user2.Unlocks()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name : "Faisal",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name : "Budi",
		Balance: 1000000, //
	}

	go Transfer(&user1, &user2, 100000)
	
	go Transfer(&user2, &user1, 200000)

	time.Sleep(time.Second * 3)
	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)
}