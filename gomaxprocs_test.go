package belajargolanggoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()	
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("TOTAL CPU", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("TOTAL THREAD", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total GOROUTINE", totalGoroutine)
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()	
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("TOTAL CPU", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("TOTAL THREAD", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total GOROUTINE", totalGoroutine)
}