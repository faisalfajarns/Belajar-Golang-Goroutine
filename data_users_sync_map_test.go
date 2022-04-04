package belajargolanggoroutine

import (
	"fmt"
	_ "reflect"
	"sync"
	"testing"
)


type Users struct {
	Name string
	Age int
}

func AddToUsersData(data *sync.Map,  users string,group *sync.WaitGroup){
	defer group.Done()
	group.Add(1)

	fmt.Println("aa",users)
	data.Store(users, users)
}

func TestMapDataUsers(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}
	users := []Users{
		{
			Name: "Faisal",
			Age: 20,
		},
		{
			Name: "Fahmi",
			Age: 20,
		},
	}


	for _, value := range users {
			fmt.Println("NILAI", value)
			go AddToUsersData(data, value.Name ,group)
	}



	// fmt.Println(users)

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println("INI KEYNYA", key)
		fmt.Println(key, " : ", value)
		return true
	})
}