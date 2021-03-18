package main

import (
	"fmt"
	"sync"
	"time"
)

func primary(counterChan chan int) {
	counter := 0

	for i := 0; i < 5; i++ {
		counterChan <- counter
		fmt.Println("Counter: ", counter)
		counter += 1
		time.Sleep(time.Second)
	}
}

func backup(counterChan chan int, wg *sync.WaitGroup) {
	var counter int
	lastMessage := time.Now()

backup:
	for {
		select {
		case x := <-counterChan:
			counter = x
			lastMessage = time.Now()
		default:
			if time.Since(lastMessage) > 3*time.Second {
				break backup
			}
		}
	}

	fmt.Println("***** BACKUP TAKEOVER *****")
	go func() {
		wg.Add(1)
		backup(counterChan, wg)
	}()

	for i := 0; i < 5; i++ {
		counter += 1
		counterChan <- counter
		fmt.Println("Counter: ", counter)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	counterChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		primary(counterChan)
		wg.Done()
	}()

	go func() {
		backup(counterChan, &wg)
	}()

	wg.Wait()
}
