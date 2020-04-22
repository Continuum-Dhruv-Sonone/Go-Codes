package sort

import "fmt"

func SelectionSort() {
	s := []int{23, 42, 35, 10, 34}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	fmt.Println(s)
}
