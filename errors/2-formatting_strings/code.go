package main

import "fmt"

func getSMSErrorString(cost float64, recipient string) string {
	// ?
   s:=fmt.Sprintf("SMS that costs %.2f to be sent to '%s' can not be sent",cost,recipient)
	return s
}

// don't edit below this lineRECIPIENT'

func test(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("====================================")
}

func main() {
	test(1.4, "+1 (435) 555 0923")
	test(2.1, "+2 (702) 555 3452")
	test(32.1, "+1 (801) 555 7456")
	test(14.4, "+1 (234) 555 6545")
}
