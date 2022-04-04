package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func ResponseSelectChannel(chn chan string){
	time.Sleep(time.Second * 2)
	chn <- "Faisal"
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go ResponseSelectChannel(channel1)
	go ResponseSelectChannel(channel2)
	counter := 0
	for {

	select {
		case data := <- channel1 :
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <- channel2 :
			fmt.Println("Data dari channel 2 ", data)
			counter++
		}
		if counter == 2 {
			break
		}
		
	}
}