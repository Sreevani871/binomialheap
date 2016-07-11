/* Double linked list is used to store values of a binomial heap */

package binomialheapQ

/* Method to insert node into the forest*/
func InsertDL(head **BinomialHeapNode, node *BinomialHeapNode) {

	var prev *BinomialHeapNode
	var next *BinomialHeapNode

	prev = nil
	next = *head
	for next != nil && node.order < next.order {
		prev = next
		next = next.rightsibling
	}
	if prev == nil && next == nil {
		*head = node
	} else if prev == nil && next != nil {
		node.rightsibling = *head
		*head = node
	} else if prev != nil && next == nil {
		prev.rightsibling = node
		node.rightsibling = next
	}
}

/* Method to remove node from forest */
func RemoveFromDL(head **BinomialHeapNode, node *BinomialHeapNode) {

	leftsib := GetLeftsibling(*head, node)
	if leftsib == nil {
		*head = node.rightsibling
	} else {
		leftsib.rightsibling = node.rightsibling
	}
	node.rightsibling = nil
}

/* To remove node form binomial heap need to find leftsibling,So this method returns leftsibling*/
func GetLeftsibling(head *BinomialHeapNode, node *BinomialHeapNode) *BinomialHeapNode {

	if head == node {
		return nil
	}
	current := head
	for current.rightsibling != node {
		current = current.rightsibling
	}
	return current
}

/* It returns node with same order of parameter node passed,otherwise returns nil */
func GetNodeWithOrder(head *BinomialHeapNode, order int) *BinomialHeapNode {
	current := head
	for current != nil {
		if current.order == order {
			return current
		}
		current = current.rightsibling
	}
	return nil
}

/* Returns minimum node from binomial tree */
func GetMinimumNode(head *BinomialHeapNode) *BinomialHeapNode {
	min := head
	current := head.rightsibling

	for current != nil {
		if current.value < min.value {
			min = current
		}
		current = current.rightsibling
	}
	return min
}

/* Iterates trough level wise */
func NodeIterator(head *BinomialHeapNode) []*BinomialHeapNode {

	var arr []*BinomialHeapNode
	trnode := head
	for trnode != nil {
		arr = append(arr, trnode)

		trnode = trnode.rightsibling
	}
	return arr
}
