package sort

func InsertSort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				for j != i {
					arr[i], arr[j] = arr[j], arr[i]
					j++
				}
			}
		}
	}
}
