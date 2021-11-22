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
	// qsortTowPointer
	qsort(s, 0, len(s)-1)
}

func quickSortTwoPointer(s []int) {
	// qsortTowPointer
	qsortTwoPointer(s, 0, len(s)-1)
}

func qsort(s []int, i, j int) {
	if i >= j {
		return
	}

	x := s[j]
	ix := i
	for k := i; k < j; k++ {
		if s[k] < x {
			s[ix], s[k] = s[k], s[ix]
			ix++
		}
	}
	s[ix], s[j] = s[j], s[ix]

	qsort(s, i, ix-1)
	qsort(s, ix+1, j)
}

func qsortTwoPointer(s []int, i, j int) {
	if i >= j {
		return
	}

	tmp := s[i : j+1]
	_ = tmp
	x := s[j]
	ix := j
	pi := i
	pj := j
	for i < j {
		for s[i] <= x && i < j {
			i++
		}

		for s[j] >= x && i < j {
			j--
		}

		if i < j {
			s[i], s[j] = s[j], s[i]
		}
	}
	s[ix], s[i] = s[i], s[ix]
	ix = i
	qsortTwoPointer(s, pi, ix-1)
	qsortTwoPointer(s, ix+1, pj)
}

func quickSortIter(s []int) {
	i := 0
	j := len(s) - 1
	x := s[0]
	ix := 0

	for i < j {
		for s[i] < x {
			i++
		}

		for s[j] > x {
			j--
		}

		if i < j {
			s[i], s[j] = s[j], s[i]
		}
	}
	s[ix], s[j] = s[j], s[ix]
}

func mergeSort(s []int) {
	doMergeSort(s, 0, len(s)-1)
}

func doMergeSort(s []int, p, q int) {
	if p >= q {
		return
	}

	x := (p + q) / 2
	doMergeSort(s, p, x)
	doMergeSort(s, x+1, q)
	merge(s, p, x, q)
}

func merge(s []int, start, mid, end int) {
	n := end - start + 1
	i := start
	j := mid + 1
	temp := make([]int, n)
	k := 0
	for i <= mid && j <= end {
		if s[i] <= s[j] {
			temp[k] = s[i]
			k++
			i++
		} else {
			temp[k] = s[j]
			k++
			j++
		}
	}

	for i <= mid {
		temp[k] = s[i]
		k++
		i++
	}

	for j <= end {
		temp[k] = s[j]
		k++
		j++
	}

	for i := 0; i < n; i++ {
		s[start+i] = temp[i]
	}
}

func heapSort(s []int) {

}
