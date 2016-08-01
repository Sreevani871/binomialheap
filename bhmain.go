/* main program for inserting elements into the binomial heap ,remvoing elements from binomial heap and merging two heaps of same type  */

package main

import S "github.com/Sreevani871/binomialheap"
import "fmt"

func main() {

	bh := S.New()
	fmt.Println("Choose ur option\n1)Insert\n2)DelMin\n3)FindMin\n4)Merge\n5)Print")
	var op int
	fmt.Scanf("%d", &op)
	for true {
		switch op {
		case 1:
			{
				var value int
				fmt.Scanf("%d", &value)
				bh.Insert(value)
			}
		case 2:
			{
				min := bh.ExtractMin()
				fmt.Println("Deleted element", min)
			}
		case 3:
			{
				min := bh.GetMinimumValue()
				fmt.Println("Minimum element", min)
			}
		case 4:
			{
				bh1 := S.New()
				bh1.Insert(3)
				bh1.Insert(2)
				bh1.Insert(1)
				bh1.Insert(0)
				bh1.Insert(-1)
				bh1.Insert(-2)
				bh1.Insert(300)
				bh1.Insert(200)
				bh1.Print()
				bh.Merge(bh1)
			}
		case 5:
			{
				bh.Print()
			}
		default:
			{

			}
		}
		fmt.Println("choose operation")
		fmt.Scanf("%d", &op)
		if op > 5 {
			break

		}
	}
}
