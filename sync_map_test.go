package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup){

	defer group.Done()
	group.Add(1)
	data.Store(value, value)
}

func TestSyncMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}


	for i:=0 ; i<=5 ; i++{
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println("KEY", key)
		fmt.Println(key, " : ", value)
		return true
	})

} 