package binomialheap

import (
	S "github.com/Sreevani871/binomialheap"
	"testing"
)

func BenchmarkInsert(b *testing.B) {
	heap := S.New()
	for j := 0; j < b.N; j++ {
		heap.Insert(5)
	}
}

func BenchmarkDel_Min(b *testing.B) {
	heap := S.New()
	heap.Insert(4)
	for i := 0; i < b.N; i++ {
		heap.ExtractMin()
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
	n2 := heap.Insert(-100)
	for i := 0; i < b.N; i++ {
		heap.Decrease_Key(n2, -200)
	}

}

func BenchmarkFindMin(b *testing.B) {
	heap1 := insert(0, 100000)
	for i := 0; i < b.N; i++ {
		heap1.GetMinimumValue()
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	heap1 := insert(0, 10)
	for i := 0; i < b.N; i++ {
		heap1.IsEmpty()
	}
}

func insert(start, end int) *S.Binomial_Heap {
	heap := S.New()

	for i := start; i < end; i++ {
		heap.Insert(i)
	}
	return heap

}
