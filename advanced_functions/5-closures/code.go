package main

import "fmt"

/*
	func adder() func(int) int {
		// ?
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

// don't touch below this line

	type emailBill struct {
		costInPennies int
	}

	func test(bills []emailBill) {
		defer fmt.Println("====================================")
		countAdder, costAdder := adder(), adder()
		for _, bill := range bills {
			fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
		}
	}

	func main() {
		test([]emailBill{
			{45},
			{32},
			{43},
			{12},
			{34},
			{54},
		})

		test([]emailBill{
			{12},
			{12},
			{976},
			{12},
			{543},
		})

		test([]emailBill{
			{743},
			{13},
			{8},
		})
	}
*/
func initSeq() func() int {

	i := 0

	return func() int {
		i++
		return i
	}

}

func main() {
	nextInt := initSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := initSeq()
	fmt.Println(newInts())

	//string1 := "stuff"
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	for index, value := range sample {
		//print value in hexa format
		fmt.Printf("index -> %v , value -> %x\n", index, value)
	}

}
