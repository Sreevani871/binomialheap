package binomialheapQ

import (
	S "github.com/Sreevani871/binomialheapQ"
	"testing"
)

func BenchmarkInsert(b *testing.B) {
	heap := S.New()
	for j := 0; j < b.N; j++ {
		for i := 0; i < 100000; i++ {
			heap.Insert(i)
		}
	}
}

func BenchmarkDel_Min(b *testing.B) {
	heap := S.New()
	heap.Insert(4)
	for i := 0; i < b.N; i++ {
		heap.Del_Min()
	}

}

func BenchmarkFind_Min(b *testing.B) {
	heap := insert(0, 100000)
	for i := 0; i < b.N; i++ {
		heap.Find_Min()
	}

}

func BenchmarkMerge(b *testing.B) {
	heap1 := insert(0, 100000)
	heap2 := insert(100000, 200000)
	for i := 0; i < b.N; i++ {
		heap1.Merge(heap2)
	}
}
func BenchmarkDecrease_Key(b *testing.B) {
	heap := insert(0, 100000)
	for i := 0; i < b.N; i++ {
		heap.Decrease_Key(40, 43)
	}

}

func insert(start, end int) *S.Binomial_Heap {
	heap := S.New()

	for i := start; i < end; i++ {
		heap.Insert(i)
	}
	return heap

}
