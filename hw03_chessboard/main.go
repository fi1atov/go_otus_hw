package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%d", &size)
	for x := 1; x <= size; x++ {
		if x%2 == 0 {
			fmt.Printf(" ")
		}
		for y := 1; y <= size-1; y++ {
			if y%2 != 0 {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		if x%2 != 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
}
