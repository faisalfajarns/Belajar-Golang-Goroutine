package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup){
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello Faisal " )

	time.Sleep(time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i:=0 ; i <= 5; i++ {
		go RunAsynchronous(group)
	}

 group.Wait()
 fmt.Println("Complete")
}