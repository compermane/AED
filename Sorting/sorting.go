package sorting

func BubbleSort(l []float64) []float64 {
	n := len(l)

	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - i - 1; j++ {
			if l[j] > l[j + 1] {
				l[j], l[j + 1] = l[j + 1], l[j]
			}
		}
	}

	return l
}

func MergeSort(l []float64) []float64 {
	if len(l) == 1 {
		return l
	}

	half := int(len(l) / 2)
	left := MergeSort(l[:half])
	right := MergeSort(l[half:])

	i, j := 0, 0
	ret := []float64{}

	for ;; {
		if i == len(left) || j == len(right) {
			break
		} else {
			if left[i] <= right[j] {
				ret = append(ret, left[i])
				i++
			} else {
				ret = append(ret, right[j])
				j++
			}
		}
	}

	for ; i < len(left); i++ {
		ret = append(ret, left[i])
	}

	for ; j < len(right); j++ {
		ret = append(ret, right[j])
	}

	return ret
}