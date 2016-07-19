package binomialheap

import (
	S "github.com/Sreevani871/binomialheap"
	"testing"
)

func BenchmarkInsertDL(b *testing.B) {
	head := S.NewBinomialHeapNode(8)
	node := S.NewBinomialHeapNode(7)
	for i := 0; i < b.N; i++ {
		S.InsertDL(&head, node)
	}
}

func BenchmarkRemoveFromDL(b *testing.B) {
	head := S.NewBinomialHeapNode(8)
	node := S.NewBinomialHeapNode(7)
	S.InsertDL(&head, node)
	node1 := S.NewBinomialHeapNode(98)
	S.InsertDL(&head, node1)
	for i := 0; i < b.N; i++ {
		S.RemoveFromDL(&head, node)
	}

}
