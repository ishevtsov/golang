package main

import "fmt"

func main() {
	for i := 0; i <= 12; i++ {
		switch i {
		case 0:
			fmt.Println(i, "Zero")
		case 1:
			fmt.Println(i, "One")
		case 2:
			fmt.Println(i, "Two")
		case 3:
			fmt.Println(i, "Three")
		case 4:
			fmt.Println(i, "Four")
		case 5:
			fmt.Println(i, "Five")
		case 6:
			fmt.Println(i, "Six")
		case 7:
			fmt.Println(i, "Seven")
		case 8:
			fmt.Println(i, "Eight")
		case 9:
			fmt.Println(i, "Nine")
		case 10:
			fmt.Println(i, "Ten")
		default:
			fmt.Println("Unknown")
		}
	}
}
