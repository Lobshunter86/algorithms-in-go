package sorts

import (
	"math/rand"
	"testing"
	"time"
)

const (
	MAXSIZE = 8192
	TIMES   = 1024
)

func arrayGen() []int {
	rand.Seed(time.Now().UnixNano())

	size := rand.Int31n(MAXSIZE)
	array := make([]int, size)

	for i := range array {
		array[i] = rand.Int()
	}
	return array
}

func maxArray() []int {
	rand.Seed(time.Now().UnixNano())
	array := make([]int, MAXSIZE)
	for i := range array {
		array[i] = rand.Int()
	}

	return array
}

func isSorted(array []int) bool {
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			return false
		}
	}
	return true
}

// actural sorting time cost = bench sort time - bench shuffle time
func BenchmarkShuffTime(B *testing.B) {
	array := maxArray()

	B.ResetTimer()
	for index := 0; index < B.N; index++ {
		rand.Shuffle(len(array), func(i, j int) { array[i], array[j] = array[j], array[i] })
	}
}

func TestHeapSort(T *testing.T) {
	for index := 0; index < TIMES; index++ {
		array := arrayGen()
		HeapSort(array)
		if !isSorted(array) {
			T.Fatal("Wrong Order")
		}
	}
}

func BenchmarkHeapSort(B *testing.B) {
	array := maxArray()

	B.ResetTimer()
	for index := 0; index < B.N; index++ {
		HeapSort(array)
		rand.Shuffle(len(array), func(i, j int) { array[i], array[j] = array[j], array[i] })
	}
}
func TestQuickSort(T *testing.T) {
	for index := 0; index < TIMES; index++ {
		array := arrayGen()
		QuickSort(array)
		if !isSorted(array) {
			T.Fatal("Wrong Order")
		}
	}
}

func BenchmarkQuickSort(B *testing.B) {
	array := maxArray()

	B.ResetTimer()
	for index := 0; index < B.N; index++ {
		QuickSort(array)
		rand.Shuffle(len(array), func(i, j int) { array[i], array[j] = array[j], array[i] })
	}
}

func TestMergeSort(T *testing.T) {
	for index := 0; index < TIMES; index++ {
		array := arrayGen()
		MergeSort(array)
		if !isSorted(array) {
			T.Fatal("Wrong Order")
		}
	}
}

func BenchmarkMergeSort(B *testing.B) {
	array := maxArray()

	B.ResetTimer()
	for index := 0; index < B.N; index++ {
		MergeSort(array)
		rand.Shuffle(len(array), func(i, j int) { array[i], array[j] = array[j], array[i] })
	}
}

func TestThreeWayQuickSort(T *testing.T) {
	for index := 0; index < TIMES; index++ {
		array := arrayGen()
		ThreeWayQuickSort(array)
		if !isSorted(array) {
			T.Fatal("Wrong Order")
		}
	}
}

func BenchmarkThreeWayQuickSort(B *testing.B) {
	array := maxArray()

	B.ResetTimer()
	for index := 0; index < B.N; index++ {
		ThreeWayQuickSort(array)
		rand.Shuffle(len(array), func(i, j int) { array[i], array[j] = array[j], array[i] })
	}
}
