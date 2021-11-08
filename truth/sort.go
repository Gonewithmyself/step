package truth

func sort1() {
	var arr []int
	for i := 0; i < 10000; i++ {
		arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
		bubbleSort(arr)
	}
}

func sort2() {
	for i := 0; i < 10000; i++ {
		arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
		bubbleSort(arr)
	}
}

func bubbleSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j+1] > arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
