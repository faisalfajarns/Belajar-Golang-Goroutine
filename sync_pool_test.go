package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := &sync.Pool{}
	
	//Menambahkan default value
	// pool := sync.Pool{
	// 	New: func() interface{} {
	// 		return "NEW"
	// 	},
	// }

	pool.Put("Faisal")
	pool.Put("Fahmi")
	pool.Put("Aryo")

	for i:= 0; i < 10; i++{
		go func() {
				data:= pool.Get()
				fmt.Println(data)
				time.Sleep(time.Second * 1)
				pool.Put(data)
				}()
	}

	time.Sleep(time.Second * 11)
	fmt.Println("Selesai")
}

