package sort

import "fmt"

func InsertionSort() {
	s := []int{23, 42, 35, 10, 34}

	for i := 1; i < len(s); i++ {
		data := s[i]
		j := i - 1

		for j >= 0 && data < s[j] {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = data
	}

	fmt.Println(s)
}
