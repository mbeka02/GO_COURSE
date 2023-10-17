package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}



func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 100)
	fmt.Println("========================")
}

func main() {
	test("Hello there Stacy!")
	test("Hi there John!")
	test("Hey there Jane!")

	text := make(chan string, 2)
	go func() {
		text <- "hello"
		text <- "world"
	}()

	fmt.Println(<-text)
	fmt.Println(<-text)

}
