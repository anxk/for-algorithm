package sort

func HeapSort(arr []int) {
	length := len(arr)
	for i := length / 2; i >= 0; i-- {
		adjust(arr, i)
	}
	for i := length - 1; i >= 1; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		adjust(arr[:i], 0)
	}
}

func adjust(arr []int, i int) {
	length := len(arr)
	left, right := 2*i+1, 2*i+2
	root := i

	if left < length && arr[left] > arr[root] {
		root = left
	}
	if right < length && arr[right] > arr[root] {
		root = right
	}

	if root != i {
		arr[i], arr[root] = arr[root], arr[i]
		adjust(arr, root)
	}
}
