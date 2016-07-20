/* Binomial heap is forest consisting of binomial trees which satisfy min heap property.
Here it consists of Binomialheap forest structure and implementation of binomial heap interface methods */

package binomialheap

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
func (bh *Binomial_Heap) Insert(value int) *BinomialHeapNode {
	bh.size += 1                          //incrementing no.of elements
	newnode := NewBinomialHeapNode(value) //newnode is struct of binomialheap node having node value
	bh.put(newnode)                       //inserting newnode into the DL
	return newnode
}

/* Method for deleting minimum element from the binomial heap */
func (bh *Binomial_Heap) ExtractMin() int {
	if bh.forest_head == nil {
		//fmt.Println("Heap is empty")
		return -1
	}
	bh.size -= 1
	min := GetMinimumNode(bh.forest_head)
	RemoveFromDL(&bh.forest_head, min)

	for _, child := range NodeIterator(min.childrenNode) {

		RemoveFromDL(&min.childrenNode, child)
		bh.put(child)
	}

	return min.Value
}

/* Find_Min method gets the minimum node value from binomial heap*/
func (bh *Binomial_Heap) GetMinimumValue() int {
	if bh.forest_head == nil {
		return -1
	}
	head := bh.forest_head
	min := GetMinimumNode(head)
	return min.Value

}
func (bh *Binomial_Heap) Decrease_Key(node *BinomialHeapNode, decval int) *BinomialHeapNode {
	resnode := node.Decrease_Key1(decval)
	return resnode
}

/* Decrease_key method decreases kenode value by passed value*/

/*func (bh *Binomial_Heap) Decrease_Key(keyvalue int, decvalue int) {

	if bh.forest_head == nil {
		fmt.Println("binomial heap is empty")
	}

	for _, node := range NodeIterator(bh.forest_head) { //Iterate through all tree nodes
		node.SearchDFS(keyvalue, decvalue)

	}

}
*/
/* SearchDFS method used to find the node which contains keyvalue to decrease */
/*func (bn *BinomialHeapNode) SearchDFS(keyvalue int, decvalue int) {
	var check int
	if bn.value == keyvalue {

		bn.Decrease_Key1(decvalue) //If keynode found Decrease_Key1 method called to decrease the keyvalue by passed decval
		check = 1
	}
	for _, child := range NodeIterator(bn.childrenNode) { //Iterating binomial heap downwards
		if check != 1 {
			child.SearchDFS(keyvalue, decvalue)
		} else {
			break
		}

	}

}*/

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

func (bh *Binomial_Heap) IsEmpty() bool {
	return bh.forest_head == nil
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
		fmt.Println("bthead", node.Value, node.order) //Printing binomial tree head nodes and order
		node.PrintDFS()                               //Iteration through each tree in binomial tree
	}
}

/* Print the nodes in binomial tree in depth first search manner*/
func (bn *BinomialHeapNode) PrintDFS() {

	fmt.Printf("Value:%d Order:%d \n", bn.Value, bn.order)

	for _, child := range NodeIterator(bn.childrenNode) { //Iterating binomial heap downwards
		fmt.Printf("Value:%d Order:%d Parent:%d\n", child.Value, child.order, child.parent.Value)
		child.PrintDFS()
	}
}

/* This method returns number of trees present in binomial heap  */
func (bh *Binomial_Heap) Trees() int {
	n := bh.size
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
func (bh *Binomial_Heap) FindMin() int {
	srnode := GetMinimumNode(bh.forest_head)
	return srnode.Value
}
