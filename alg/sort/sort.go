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
