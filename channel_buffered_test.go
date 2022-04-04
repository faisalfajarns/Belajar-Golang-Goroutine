package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	
	go func(){
		channel <- "Faisal"
		channel <-"Fajar"
		channel <-"Nursaid"
		close(channel)
		}()
		
		// go func ()  {			
			// 	fmt.Println(<- channel)
			// 	fmt.Println(<- channel)
			// 	fmt.Println(<- channel)
			// 		}()
			for data := range channel{
				fmt.Println(data)
			}
	time.Sleep(time.Second *2)
	fmt.Println("Data Selesai")
}