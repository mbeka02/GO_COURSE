package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomeStuff(t time.Duration, wg *sync.WaitGroup) {

	fmt.Println("doing some stuff ....")
	time.Sleep(t)
	fmt.Println("done with the task....")
	defer wg.Done()
}

func main() {
	startTime := time.Now()
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go doSomeStuff(time.Second*4, wg)

	go doSomeStuff(time.Second*8, wg)
	wg.Wait()
	fmt.Printf("The elapsed time is %v \n", time.Since(startTime))

}
