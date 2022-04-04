package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func ResponseSelectDefaultChannel(chn chan string){
	time.Sleep(time.Second * 2)
	chn <- "Faisal"
}

func TestSelectDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go ResponseSelectDefaultChannel(channel1)
	go ResponseSelectDefaultChannel(channel2)
	counter := 0
	for {

	select {
		case data := <- channel1 :
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <- channel2 :
			fmt.Println("Data dari channel 2 ", data)
			counter++

		default :
			fmt.Println("Menunggu Data...")
		}
		if counter == 2 {
			break
		}
		
	}
}