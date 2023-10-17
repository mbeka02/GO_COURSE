package main

import (
	"fmt"
	"time"
)

func waitForDbs(numDBs int, dbChan chan struct{}) {
	// block until we receive numDBs tokens
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}

}

func test(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
	fmt.Println("All databases are online!")
	fmt.Println("=====================================")
}

func main() {
	test(30)
	test(40)
	test(5555)
}

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()
	return ch
}
