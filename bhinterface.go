/* Binomial Heap interface methods to insert elements,remove elements,no.of elements in binomial heap,no.of trees in binomial heap and merge two same type binomial heaps*/

package binomialheap

type BinomialHeap interface {
	Insert(int) *BinomialHeapNode
	ExtractMin() int
	GetMinimumValue()int
	Decrease_Key()
	Merge(BinomialHeap)
}
