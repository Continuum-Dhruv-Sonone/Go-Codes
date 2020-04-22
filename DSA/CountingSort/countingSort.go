package main

import "fmt"

func main() {

	arr := []int{2, 1, 3, 6, 4, 5, 1, 2}

	count := [10]int{}
	output := make([]int, len(arr))

	for _, val := range arr {
		index := val % 10
		count[index] = count[index] + 1
	}

	sum := 0

	for index, val := range count {
		sum = sum + val
		count[index] = sum
	}

	for _, val := range arr {
		output[count[val]-1] = val
		count[val] = count[val] - 1
	}

	fmt.Println(output)
}
