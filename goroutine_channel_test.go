package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

type Message struct {
	data chan string
}

func (message *Message) FirstGoroutine() {
	newmsg := &message.data
	*newmsg <- "hi"
	go SecondGoroutine(*newmsg)


	close(*newmsg)
}

func SecondGoroutine(msg chan string) {

	ambil := <- msg
	fmt.Println(ambil)
}

func TestA(t *testing.T){
	chn := make(chan string)
	msg := &Message{
		data: chn,
	}

	msg.FirstGoroutine()
}

func Sender(chndikasihdata chan string){
	chndikasihdata <- "hey channel"
}

// func Receiver(chn chan string){
// 	defer close(chn)
// 	fmt.Println("result", <- chn)
// }

func TestSimpleGoroutine(t *testing.T){
	
	chnngambil := make(chan string)

	go Sender(chnngambil)
	// go Receiver(chn)

	fmt.Println(<-chnngambil)

	time.Sleep(time.Second)

}