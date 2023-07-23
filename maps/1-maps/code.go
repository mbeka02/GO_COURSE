package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	// create map
	m := make(map[string]user)

	//ensure names and phonenumber array are of equal length
	if len(phoneNumbers) != len(names) {
		return nil, errors.New("invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		name := names[i]
		m[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[i],
		}
	}
	

	return m, nil

}



type user struct {
	name        string
	phoneNumber int
}

func test(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("====================================")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)
	}
}

func main() {
	test(
		[]string{"John", "Bob", "Jill"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"John", "Bob"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"George", "Sally", "Rich", "Sue"},
		[]int{20955559812, 38385550982, 48265554567, 16045559873},
	)
}
