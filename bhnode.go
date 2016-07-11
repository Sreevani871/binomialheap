/* This program represents structure of node in binomial heap and implements union operation which is essential in performing any operations with binomial heap */
package binomialheapQ

//import "fmt"

/* BinominalHeapNode struct declaration*/
type BinomialHeapNode struct {
	value        int               //value for storig element into the bh
	order        int               //order represents order of node means number of children it has
	parent       *BinomialHeapNode //this is a parent node
	childrenNode *BinomialHeapNode //childrenNode used to link subtrees whenever union operation performed on binomialheap
	rightsibling *BinomialHeapNode //rightsibling represents

}

/* Returns a struct of node,assigning value to it*/
func NewBinomialHeapNode(value int) *BinomialHeapNode {
	return &BinomialHeapNode{
		value:        value,
		order:        0,
		parent:       nil,
		childrenNode: nil,
		rightsibling: nil,
	}
}

/* contain iserts the linked nodes into into the DL */
func (bn *BinomialHeapNode) link(other *BinomialHeapNode) {

	InsertDL(&bn.childrenNode, other)
	other.parent = bn //make bn node as the parent node of other node
}

/* Union operation takes two binomialheapnodes of same order and compare its parent values and make the one as a child of other which has less parent value*/
func Union(n1 *BinomialHeapNode, n2 *BinomialHeapNode) *BinomialHeapNode {

	if n1.value < n2.value {
		n1.order += 1
		n1.link(n2)
		return n1
	} else {
		n2.order += 1
		n2.link(n1)
		return n2
	}
}
