package sorts

// O(n) space
func MergeSort(array []int) {
	buf := append([]int(nil), array...)
	doMergeSort(buf, array)
}

func doMergeSort(array []int, buf []int) { // when function return, buf[] is sorted
	if len(array) < 2 {
		return
	}

	doMergeSort(buf[:len(array)/2], array[:len(array)/2])
	doMergeSort(buf[len(array)/2:], array[len(array)/2:])

	// merge
	left := array[:len(array)/2]
	right := array[len(array)/2:]
	indexL, indexR, i := 0, 0, 0

	for ; indexL < len(left) && indexR < len(right); i++ {
		if left[indexL] < right[indexR] {
			buf[i] = left[indexL]
			indexL++
		} else {
			buf[i] = right[indexR]
			indexR++
		}
	}

	if indexL < len(left) {
		for i < len(buf) {
			buf[i] = left[indexL]
			indexL++
			i++
		}
	} else {
		for i < len(buf) {
			buf[i] = right[indexR]
			indexR++
			i++
		}
	}
}

// naive version, easier to understand, too many memory allocations, unfriendly to GC, worse performance
// unlike theoretical analysis, this version will acturally allocate O(nlogn) space, it's really slow
// func MergeSort(array []int) []int {
// 	if len(array) < 2 {
// 		return array
// 	}

// 	left := MergeSort(array[:len(array)/2])
// 	right := MergeSort(array[len(array)/2:])

// 	/// merge
// 	indexL, indexR, i := 0, 0, 0
// 	result := make([]int, len(array))
// 	for ; indexL < len(left) && indexR < len(right); i++ {
// 		if left[indexL] < right[indexR] {
// 			result[i] = left[indexL]
// 			indexL++
// 		} else {
// 			result[i] = right[indexR]
// 			indexR++
// 		}
// 	}
// 	if indexL < len(left) {  // for simplicity
// 		result = append(result[:i], left[indexL:]...)
// 	} else {
// 		result = append(result[:i], right[indexR:]...)
// 	}
// 	return result
// }
