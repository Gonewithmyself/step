package sort

func numK(data []int32, k int) int32 {
	pidx := len(data) - 1
	pivot := data[pidx]

	for i := pidx - 1; i >= 0; i-- {
		if data[i] < pivot {
			data[pidx] = data[i]
			// data[i] = pivot
			// pivot = data[pidx]
			pidx = i
		}
	}
	data[pidx] = pivot

	return int32(pidx)
}
