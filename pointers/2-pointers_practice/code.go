package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	// ?
	msg := *message
	msg = strings.ReplaceAll(msg, "dang", "****")
	msg = strings.ReplaceAll(msg, "shoot", "*****")
	msg = strings.ReplaceAll(msg, "heck", "****")

	*message = msg

}

// don't touch below this line

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	messages1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	test(messages1)
	test(messages2)

	var x int = 50
	var y *int = &x
	*y = 100
	fmt.Println(*y)
	fmt.Println(y)
	fmt.Println(&y)
	fmt.Println(&x)
	
}
