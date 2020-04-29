package main

import (
	"fmt"
)

func main() {
	power := 9000
	fmt.Printf("default power: %d\n", power)

	name, power := "Goku", 1000
	fmt.Printf("%s's power is over %d\n", name, power)
}
