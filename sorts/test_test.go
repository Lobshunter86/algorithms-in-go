package sorts

import (
	"math/rand"
	"testing"
	"time"
)

const (
	MAXSIZE = 1024
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

func isSorted(array []int) bool {
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			return false
		}
	}
	return true
}
func TestQuickSort(T *testing.T) {
	for index := 0; index < TIMES; index++ {
		array := arrayGen()
		quickSort(array)
		if !isSorted(array) {
			T.Fatal("Wrong Order")
		}
	}
}

func BenchmarkQuickSort(B *testing.B) {
	array := make([]int, MAXSIZE)
	for i := range array {
		array[i] = rand.Int()
	}

	B.ResetTimer()
	for index := 0; index < B.N; index++ {
		quickSort(array)
	}
}

func TestMergeSort(T *testing.T) {
	for index := 0; index < TIMES; index++ {
		array := arrayGen()
		mergeSort(array)
		if !isSorted(array) {
			T.Fatal("Wrong Order")
		}
	}
}

func BenchmarkMergeSort(B *testing.B) {
	array := make([]int, MAXSIZE)
	for i := range array {
		array[i] = rand.Int()
	}

	B.ResetTimer()
	for index := 0; index < B.N; index++ {
		mergeSort(array)
	}
}
