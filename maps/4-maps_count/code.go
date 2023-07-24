package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func getCounts(userIDs []string) map[string]int {

//create map
	m := map[string]int{}

	for _, id := range userIDs {
		//if key does not exist
		if _, ok := m[id]; !ok {
			//init value to 0 for non-existent key
			m[id] = 0
		}
		//tracks no .of times string/key has showed up and incs that count
		m[id]++
	}
	return m
}



func test(userIDs []string, ids []string) {
	fmt.Printf("Generating counts for %v user IDs...\n", len(userIDs))

	counts := getCounts(userIDs)
	fmt.Println("Counts from select IDs:")
	for _, k := range ids {
		v := counts[k]
		fmt.Printf(" - %s: %d\n", k, v)
	}
	fmt.Println("=====================================")
}

func main() {
	userIDs := []string{}
	for i := 0; i < 10000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprint(i))
		key := fmt.Sprintf("%x", h.Sum(nil))
		userIDs = append(userIDs, key[:2])
	}

	test(userIDs, []string{"00", "ff", "dd"})
	test(userIDs, []string{"aa", "12", "32"})
	test(userIDs, []string{"bb", "33"})
}
