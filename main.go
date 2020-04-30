package main

import "fmt"

func main() {
	lookup := make(map[string]int, 100)
	lookup["goku"] = 9000
	power, exists := lookup["vegeda"]
	total := len(lookup)
	fmt.Println(power, exists)
	fmt.Println(total)
}
