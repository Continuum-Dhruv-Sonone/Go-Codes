package sort

import "fmt"

func QuickSortCode() {

	s := []int{23, 42, 35, 10, 33}

	QuickSort(s, 0, len(s)-1)

	fmt.Println(s)
}

func QuickSort(s []int, start int, end int) {
	if start < end {
		p := pivot(s, start, end)
		QuickSort(s, start, p-1)
		QuickSort(s, p+1, end)

	}
}

func pivot(s []int, start int, end int) int {
	element := s[end]
	i := start - 1

	for j := start; j < end; j++ {
		if s[j] < element {
			i = i + 1
			s[i], s[j] = s[j], s[i]
		}
	}
	i = i + 1
	s[i], s[end] = s[end], s[i]

	return i
}
