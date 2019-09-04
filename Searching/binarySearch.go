package searching

import "fmt"

func BinarySearchCode() {
	if BinarySearch() {
		fmt.Println("Element found")
	} else {
		fmt.Println("Element not found")
	}
}

func BinarySearch() bool {
	s := []int{10, 23, 34, 35, 42}
	element := 34

	low := 0
	high := len(s) - 1

	for low <= high {
		mid := (low + high) / 2

		if s[mid] == element {
			return true
		} else if element < s[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}
