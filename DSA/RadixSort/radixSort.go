package main

import "fmt"

func main() {
	arr := []int{170, 45, 75, 90, 802, 24, 2, 66}
	radixSort(arr, len(arr))

	fmt.Println(arr)
}

func radixSort(arr []int, n int) {

	max := getMax(arr)

	for places := 1; max/places > 0; places = places * 10 {
		countingSort(arr, n, places)
	}
}

func countingSort(arr []int, n, place int) {
	count := [10]int{}
	output := make([]int, n+1)

	for i := 0; i < n; i++ {
		count[(arr[i]/place)%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] = count[i] + count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		output[count[(arr[i]/place)%10]-1] = arr[i]
		count[(arr[i]/place)%10]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}

}

func getMax(arr []int) int {
	max := arr[0]

	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	fmt.Println(max)
	return max
}
