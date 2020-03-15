package sort

func ShellSort(arr []int) {
	length := len(arr)
	for one := length / 2; one > 0; one /= 2 { // one is the gap
		for i := 0; i < one; i++ {
			for j := one + i; j < length; j += one { // insert sort
				for k := i; k < j+one; k += one {
					if arr[k] > arr[j] {
						for j != k {
							arr[j], arr[k] = arr[k], arr[j]
							k += one
						}
					}
				}
			}
		}
	}
}
