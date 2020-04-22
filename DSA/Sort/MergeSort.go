package sort

import "fmt"

func MergeSortCode() {
	s := []int{23, 42, 35, 10, 34}

	MergeSort(s, 0, len(s)-1)

	fmt.Println(s)
}

func MergeSort(s []int, first int, last int) {
	if first < last {
		mid := (first + last) / 2

		MergeSort(s, first, mid)
		MergeSort(s, mid+1, last)
		Merge(s, first, mid, last)
	}
}

func Merge(s []int, first int, mid int, last int) {

	l := []int{}
	m := []int{}
	k := 0
	for i := first; i <= mid; i++ {
		l = append(l, s[i])
	}
	k = 0
	for i := mid + 1; i <= last; i++ {
		m = append(m, s[i])
	}

	a := mid - first + 1
	b := last - mid

	i, j := 0, 0
	k = first
	for i < a && j < b {
		if l[i] > m[j] {
			s[k] = m[j]
			j++
		} else {
			s[k] = l[i]
			i++
		}
		k++
	}

	for i < a {
		s[k] = l[i]
		k++
		i++
	}

	for j < b {
		s[k] = m[j]
		k++
		j++
	}
}
