package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
channel := make(chan string)

// mengirim data ke channel
// channel <- "Faisal"

//menerima data dari channel dan disimpan kedalam variable
// data := <- channel

go func(){
time.Sleep(2 * time.Second)
channel <- "Faisal Fajar Nursaid"
fmt.Println("Selesai Mengirim Data Ke Channel")
}()
// defer close(channel)
data := <- channel

fmt.Println(data)
time.Sleep(5 * time.Second)
}