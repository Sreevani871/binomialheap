package binomialheap

import (
	S "github.com/Sreevani871/binomialheap"
	"testing"
)

func BenchmarkNewBinomialHeapNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		S.NewBinomialHeapNode(5)
	}

}

func BenchmarkLink(b *testing.B) {
	// Bhn1 S.NewBinomialHeaoNode
	Bhn1 := S.NewBinomialHeapNode(4)
	Bhn2 := S.NewBinomialHeapNode(29)
	for i := 0; i < b.N; i++ {
		Bhn1.Link(Bhn2)
	}

}

func BenchmarkUnion(b *testing.B) {
	Bhn1 := S.NewBinomialHeapNode(4)
	Bhn2 := S.NewBinomialHeapNode(29)
	for i := 0; i < b.N; i++ {
		S.Union(Bhn1, Bhn2)
	}

}

func BenchmarkDecrease_Key1(b *testing.B) {
	Bhn1 := S.NewBinomialHeapNode(13)
	for i := 0; i < b.N; i++ {
		Bhn1.Decrease_Key1(9)
	}

}
