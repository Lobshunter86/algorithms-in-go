package sorts

// performs better than when many elements are the same
// can observe better performance by modifying maxArray() in test_test.go
func ThreeWayQuickSort(array []int) {
	if len(array) < 2 {
		return
	}
	pivot := array[0]
	low, high := 0, len(array)-1

	for i := 1; i <= high; {
		if array[i] > pivot {
			array[i], array[high] = array[high], array[i]
			high--
		} else if array[i] < pivot {
			array[i], array[low] = array[low], array[i]
			i++
			low++
		} else {
			i++
		}
	}

	if high < len(array)-1 {
		ThreeWayQuickSort(array[high+1:])
	}
	if low > 0 {
		ThreeWayQuickSort(array[:low])
	}
}
