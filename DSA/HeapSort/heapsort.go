package heapsort

func heapify(arr []int, n, current int) {
	right := 2*current + 2
	left := 2*current + 1
	root := current

	if left < n && arr[left] > arr[root] {
		root = left
	}

	if right < n && arr[right] > arr[root] {
		root = right
	}

	if root != current {
		swap(arr, current, root)
		heapify(arr, n, root)
	}
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

//HeapSort sorts the given array
func HeapSort(arr []int) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		swap(arr, 0, i)
		heapify(arr, i, 0)
	}
}
