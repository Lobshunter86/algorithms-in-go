package sorts

// root = 0
// last parent of tree = (len-1) / N
// first child of node = i*N + 1
// parent of node = (i-1) / N
//

// n-ary heap is more cache friendly than binary heap, like, 5% faster with int64, 8192 eletments slice
// cpu alignment of my CPU is 64 Bytes, I think most CPUs do the same

// heap itself is not perfect, root node takes one position, so that ruin the memory alignment
// I thought about optimize the heap structure for better alignment, but it seems like a tiring job,
// Maybe I will try it in the future.

const N = 4

func HeapSort(array []int) {
	i := (len(array) - 1) / N
	for ; i >= 0; i-- {
		heapify(array, i)
	}

	for i = len(array) - 1; i > 0; i-- {
		array[i], array[0] = array[0], array[i]
		heapify(array[:i], 0)
	}
}

func heapify(array []int, start int) {
	parent := start
	child := parent*N + 1

	var maxChild int
	for child <= len(array)-1 {
		/// find max child index
		maxChild = child
		if child+N > len(array) {
			for index := child; index < len(array); index++ {
				if array[index] > array[maxChild] {
					maxChild = index
				}
			}
		} else {
			for index := child; index < child+N; index++ {
				if array[index] > array[maxChild] {
					maxChild = index
				}
			}
		}
		child = maxChild

		if array[parent] > array[child] {
			return
		} else {
			array[child], array[parent] = array[parent], array[child]
			parent = child
			child = parent*N + 1
		}
	}
}
