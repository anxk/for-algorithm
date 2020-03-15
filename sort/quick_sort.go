package sort

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	refIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[refIndex] {
			for j := refIndex; j != i; {
				arr[j], arr[i] = arr[i], arr[j]
				j++
			}
			refIndex++
		}
	}

	QuickSort(arr[0 : refIndex+1])
	QuickSort(arr[refIndex+1:])
}
