package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

//ATomic lebih cocok digunakan untuk tipe data yang primitive sedangkan mutex lebih cocok digunakan untuk tipe data struct
func TestSyncAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i<=1000 ; i++ {
		go func ()  {
			group.Add(1)
			for j:= 1; j <= 100; j++ {
				// x = x + 1
				atomic.AddInt64(&x, 1)
			}	
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Counter = ", x)
}