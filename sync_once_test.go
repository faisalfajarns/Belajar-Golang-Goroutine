package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

var coun = 0

func OnlyOnce() {
	coun++
}
//notes sync.Once ini hanya dapat digunakan kepada fungsi yang tidak memiliki paramater
func TestOnlyOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i:= 0 ; i < 100; i++{
		group.Add(1)	
		fmt.Println("a")
		go func ()  {
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Count", coun)
}