package main

import "fmt"

func binarySearch(data []int, target int) int {
	leftElem := 0
	rightElem := len(data) - 1

	for leftElem <= rightElem {
		middleElem := (leftElem + rightElem) / 2
		middleValue := data[middleElem]

		switch {
		case middleValue == target:
			return middleElem
		case middleValue < target:
			leftElem = middleElem + 1
		case middleValue > target:
			rightElem = middleElem - 1
		}
	}

	return -1
}

func main() {
	data := []int{1, 5, 12, 18, 20, 26, 46, 78, 82}

	fmt.Println(binarySearch(data, 100))
}
