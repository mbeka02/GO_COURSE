package main

import (
	"fmt"
	"time"
)

/*`sync.WaitGroup` is essentially a counter that we can
increase (to indicate we want to wait for things),
and decrease (to indicate things are done).
Then we can tell code to wait until the WaitGroup counter
reaches zero, which would mean all things have finished.*/

//var wg sync.WaitGroup

func main() {
	/*fmt.Println("start")
	wg.Add(1)
	go doStuff()
	fmt.Println("end")
	for k := 1; k < 12; k++ {
		wg.Add(1)
		go timestable(k)
	}

	wg.Wait() //block until all jobs complete*/
	ores := []string{"rocks", "ore", "ore", "ore",  "rocks"}
	oreChan := make(chan string)
	minedOreChan := make(chan string)

	//finder
	go func(o []string) {

		for _, item := range o {
			if item == "ore" {
				//send to channel
				oreChan <- item
			}

		}

	}(ores)

	//miner
	go func() {
		for i := 0; i < 3; i++ {
			// receive value from channel
			receivedOre := <-oreChan
			fmt.Printf(" Received the %s from the  finder \n", receivedOre)
			minedOreChan <- "I've mined stuff"

		}

	}()

	//smelter
	go func() {
		minerResponse := <-minedOreChan
		for i := 0; i < 3; i++ {
			fmt.Printf("From Miner : %s \n",minerResponse)
			fmt.Println("From the smelter : I'm smelting ")
		}

	}()
	//I'm assuming this just blocks and  allows the go-routines to run in the background before the program exits
	<-time.After(time.Second * 5)

}

/*func doStuff() {
	fmt.Println("...doing  some stuff")
	wg.Done() //indicate that job is done and dec by 1
}*/
