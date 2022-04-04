package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	
fmt.Println("Hai Faisal")
}

//proses async dapat dilakukan menggunakan go
//goroutine tidak cocok dilakukan untuk fungsi yang mengembalikan nilai 
func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
fmt.Println("Ups")

	time.Sleep(1*time.Second)
}

func DisplayNumber(number int){
	fmt.Println("Display : ", number)
}

func TestDisplayGoroutine(t *testing.T){
	for i := 0; i < 100000; i++ {
	go	DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}