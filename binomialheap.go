/* Binomial heap is forest consisting of binomial trees which satisfy min heap property.
Here it consists of Binomialheap forest structure and implementation of binomial heap interface methods */

package binomialheapQ

import "fmt"

/* Structure representing binomial heap forest*/
type Binomial_Heap struct {
	forest_head *BinomialHeapNode //forest_head is the leftmost node in the binomial heap
	size        int               //size represents the number of nodes in the binomial heap
}

/* New method returns new Binmomial heap whenever it calls */
func New() *Binomial_Heap {

	return &Binomial_Heap{
		forest_head: nil, size: 0,
	}
}

/* Method for inserting nodes into the binomial heap*/
func (bh *Binomial_Heap) Insert(value int) {
	bh.size += 1                          //incrementing no.of elements
	newnode := NewBinomialHeapNode(value) //newnode is struct of binomialheap node having node value
	bh.put(newnode)                       //inserting newnode into the DL
}

/* Method for deleting minimum element from the binomial heap */
func (bh *Binomial_Heap) Del_Min() int {

	bh.size -= 1
	min := GetMinimumNode(bh.forest_head)
	RemoveFromDL(&bh.forest_head, min)

	for _, child := range NodeIterator(min.childrenNode) {
		RemoveFromDL(&min.childrenNode, child)
		bh.put(child)
	}

	return min.value
}

/* Find_Min method gets the minimum node value from binomial heap*/
func (bh *Binomial_Heap) Find_Min() int {
	head := bh.forest_head
	min := GetMinimumNode(head)
	return min.value

}

/* Decrease_key method decreases kenode value by passed value*/

func (bh *Binomial_Heap) Decrease_Key(keyvalue int, decvalue int) {
	bh.Decrease(keyvalue, decvalue)

}

func (bh *Binomial_Heap) Decrease(keyvalue int, decvalue int) {

	if bh.forest_head == nil {
		fmt.Println("Empty")
	}

	for _, node := range NodeIterator(bh.forest_head) { //Iterate through all tree nodes
		node.SearchDFS(keyvalue, decvalue)

	}

}

/* SearchDFS method used to find the node which contains keyvalue to decrease */
func (bn *BinomialHeapNode) SearchDFS(keyvalue int, decvalue int) {
	if bn.value == keyvalue {

		bn.Decrease_Key1(decvalue) //If keynode found Decrease_Key1 method called to decrease the keyvalue by passed decval

	}
	for _, child := range NodeIterator(bn.childrenNode) { //Iterating binomial heap downwards

		child.SearchDFS(keyvalue, decvalue)
	}

}

/* Method for merging two binomial heaps */
func (bh *Binomial_Heap) Merge(other *Binomial_Heap) {
	bh.size += other.size

	for _, child := range NodeIterator(other.forest_head) {
		RemoveFromDL(&other.forest_head, child)
		bh.put(child)
	}
}

/* Returns the no.of nodes in binomial heap */

func (bh *Binomial_Heap) Size() int {
	return bh.size
}

/* Inserting newnode into the DL maintaing unique order trees in binomial heap */
func (bh *Binomial_Heap) put(newnode *BinomialHeapNode) {

	srnode := GetNodeWithOrder(bh.forest_head, newnode.order) //It returns whether the binomial heap contains any tree with order same as newnode
	if srnode == nil {                                        //No tree having the same order of newnode
		InsertDL(&bh.forest_head, newnode) //Inserting elements into the DL
	} else {
		RemoveFromDL(&bh.forest_head, srnode) //If any tree order in bh matches with newnode order,remove it from DL
		linkednode := Union(srnode, newnode)  //then perform union operation
		bh.put(linkednode)                    //then inset the linked node into the DL
	}
}

/* This method prints all tree nodes in binomial heap */
func (bh *Binomial_Heap) Print() {

	if bh.forest_head == nil {
		fmt.Println("binomial heap is empty")
	}
	for _, node := range NodeIterator(bh.forest_head) { //Iterate through all tree nodes
		fmt.Println("bthead", node.value, node.order) //Printing binomial tree head nodes and order
		node.PrintDFS()                               //Iteration through each tree in binomial tree
	}
}

/* Print the nodes in binomial tree in depth first search manner*/
func (bn *BinomialHeapNode) PrintDFS() {

	fmt.Printf("Value:%d Order:%d \n", bn.value, bn.order)

	for _, child := range NodeIterator(bn.childrenNode) { //Iterating binomial heap downwards
		fmt.Printf("Value:%d Order:%d Parent:%d\n", child.value, child.order, child.parent.value)
		child.PrintDFS()
	}
}

/* This method returns number of trees present in binomial heap  */
func (bh *Binomial_Heap) Trees() int {
	n := bh.size
	fmt.Println("size", n)
	var count int
	for n > 0 {
		if n%2 == 1 {
			count += 1
		}
		n = n / 2
	}
	fmt.Println(count)
	return count
}
