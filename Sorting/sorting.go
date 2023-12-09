package sorting


func QuickSort(l []float64, low int, high int) {
	if low <= high {
		p := Partition(l, low, high)
		QuickSort(l, low, p - 1)
		QuickSort(l, p + 1, high)
	}
}

func Partition(l []float64, low int, high int) int {
	pivot := l[high]
	i := low - 1

	for j := low; j <= high; j++ {
		if l[j] < pivot {
			i++
			l[i], l[j] = l[j], l[i]
		}
	}

	l[i + 1], l[high] = l[high], l[i + 1]
	return i + 1
}

func HeapSort(l []float64) {
	n := len(l)

	for i := (n / 2) - 1; i >= 0; i-- {
		Heapify(l, n, i)
	}

	for i := n - 1; i > 0; i-- {
		l[0], l[i] = l[i], l[0]
		Heapify(l, i, 0)
	}
}

func Heapify(l []float64, n int, i int) {
	largest := i
	left := (2 * i) + 1
	right := (2 * i) + 2

	if left < n && l[largest] < l[left] {
		largest = left
	}

	if right < n && l[largest] < l[right] {
		largest = right
	}

	if largest != i {
		l[i], l[largest] = l[largest], l[i]
		Heapify(l, n, largest)
	}
}

func InsertionSort(l []float64) []float64 {
	n := len(l)

	for i := 1; i < n; i++ {
		chave := l[i]
		j := i - 1

		for j >= 0 && l[j] > chave {
			l[j + 1] = l[j]
			j--
		}
		l[j + 1] = chave
	}

	return l
}

func SelectionSort(l []float64) []float64 {
	n := len(l)

	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			if l[i] > l[j] && i != j {
				l[i], l[j] = l[j], l[i]
			}
		}
	}

	return l
}

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