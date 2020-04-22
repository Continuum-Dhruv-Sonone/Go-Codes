package sort

import "fmt"

var s []int = []int{23, 42, 35, 10, 34}

func BubbleSort() {
	//s := []int{23, 42, 35, 10, 34}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j+1], s[j] = s[j], s[j+1]
			}
		}

	}

	fmt.Println(s)
}
