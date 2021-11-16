package sort

func insertSort(s []int) {
	for i := 1; i < len(s); i++ {
		for j := i; j >= 1; j-- {
			temp := s[j]
			if temp < s[j-1] {
				s[j] = s[j-1]
			} else {
				break
			}
		}
	}
}

func bubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		maxIdx := len(s) - i - 1
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[maxIdx] {
				maxIdx = j
			}
		}

		if maxIdx != len(s)-i-1 {
			s[maxIdx], s[len(s)-i-1] = s[len(s)-i-1], s[maxIdx]
		}
	}
}

func quickSort(s []int) {

}
