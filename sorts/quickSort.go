package sorts

// O(logN) space
func QuickSort(array []int) {
	if len(array) < 2 {
		return
	}

	pivot := array[0]
	low, high := 0, len(array)-1

	for i := 1; i <= high; {
		if array[i] > pivot {
			array[i], array[high] = array[high], array[i]
			high--
		} else {
			array[i], array[low] = array[low], array[i]
			i++
			low++
		}
	}

	if high < len(array)-1 {
		QuickSort(array[high+1:])
	}
	if low > 0 {
		QuickSort(array[:low])
	}
}
