/* Binomial Heap interface methods to insert elements,remove elements,no.of elements in binomial heap,no.of trees in binomial heap and merge two same type binomial heaps*/

package binomialheapQ

type BinomialHeap interface {
	Insert(int)
	Remove() int
	Size() int
	Trees() int
	Merge(BinomialHeap)
}
