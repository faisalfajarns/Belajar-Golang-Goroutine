package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int){
	defer group.Done()
	// group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("DONE",  value)
	cond.L.Unlock()
}

func TestSyncCond(t *testing.T) {

	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			cond.Signal()
		}
	}()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()
}