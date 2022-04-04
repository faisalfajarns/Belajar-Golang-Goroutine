package belajargolanggoroutine

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string) {
	
	channel <- "Faisal"
	time.Sleep(time.Second * 2)
}

func FungsiChannel(fungsi func (chan string), chn chan string){
	fungsi(chn)
}
func TestChannelAsParams(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	// go FungsiChannel(GiveMeResponse, channel)
	go GiveMeResponse(channel)

	data := <- channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

type MapAnys []map[string]interface{}
type ChnAny chan []map[string]interface{}

func GimmeInterface(data ChnAny){
	var val = MapAnys{
		{
			"Name":"faisal",
			"Age":12,
		},
		{
			"Name":"fahmi",
			"Age":15,
		},
	} 

	data <- val
	
}

func TestMap(t *testing.T) {
	chn := make(chan []map[string]interface{})

	go GimmeInterface(chn)

	for _, v := range <- chn {
		fmt.Println(reflect.TypeOf(v["Age"]))
	}
}

func OnlyIn(channel chan<- string){

	channel <- "FaisalFajar"
	time.Sleep(time.Second * 2)
}

func OnlyOut(channel <-chan string){
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)
	
	go OnlyIn(channel)
	go OnlyOut(channel)
	
	time.Sleep(time.Second * 2)
}