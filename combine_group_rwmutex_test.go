package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBankBalance struct {
	sync.Mutex
	Name string
	Balance int
}

func (acc *UserBankBalance) Lock(){
	acc.Mutex.Lock()
}

func (acc *UserBankBalance) Unlock(){
	acc.Mutex.Unlock()
}

func (account *UserBankBalance) GetBalanceAcc() int{
	account.Mutex.Lock()
	balance := account.Balance
	account.Mutex.Unlock()

	return balance
}

type Params struct {
	User1 *UserBankBalance
	User2 *UserBankBalance
	amount int
	*sync.WaitGroup
	sync.Mutex
}

func (account *UserBankBalance) ChangeBalanceAcc(amount int){
	account.Lock()
	fmt.Println("acc ",account.Name, account.Balance, amount)
	account.Balance = account.Balance + amount
	account.Unlock()
}

func TransferBalanceAcc(p *Params){
	defer p.WaitGroup.Done()

	p.WaitGroup.Add(1)
	

	// account1.Lock()
	p.Mutex.Lock()
	fmt.Println("proses 1", p.User1.Name, -p.amount)
	p.User1.ChangeBalanceAcc(-p.amount)
	// time.Sleep(time.Second *2)
	// account1.Unlock()
	p.Mutex.Unlock()
	
	// account2.Lock()
	// p.Mutex.Lock()
	fmt.Println("proses 2", p.User2.Name, p.amount)
p.User2.ChangeBalanceAcc(p.amount)
	// time.Sleep(time.Second *2)
	// p.Mutex.Unlock()
	// account2.Unlock()
	
}

func TestCombineWaitGroupRWMutex(t *testing.T) {
	group := &sync.WaitGroup{}

	acc1 := &UserBankBalance{
		Name: "Faisal",
		Balance: 10000,
	}

	acc2 := &UserBankBalance{
		Name: "Fahmi",
		Balance: 5000,
	}

	// mutex.Lock()
	go TransferBalanceAcc(&Params{
		User1: acc1,
		User2: acc2,
		amount: 2000,
		WaitGroup: group,
	})
	go TransferBalanceAcc(&Params{
		User1: acc2,
		User2: acc1,
		amount: 500,
		WaitGroup: group,
	})
	time.Sleep(time.Millisecond * 1)
	// mutex.Unlock()

	group.Wait()
	
	// fmt.Println("User ", acc1.Name, ", Balance ", acc1.Balance)
	
	fmt.Println("User ", acc1.Name, ", Balance ", acc1.Balance)
	fmt.Println("User ", acc2.Name, ", Balance ", acc2.Balance)

	
}

