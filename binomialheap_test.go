package binomialheap

import (
	"fmt"
	S "github.com/Sreevani871/binomialheap"
	"testing"
)

func Test_Insert(t *testing.T) {
	heap := S.New()

	heap.Insert(1)
	s := heap.Size()
	fmt.Println("size", s)
	if s != 1 {
		t.Errorf("something went wrong.")
	}
	v := heap.Del_Min()
	if v != 1 {
		t.Error(
			"Expected", 1,
			"got", v,
		)
	}
}

func Test_Insert_1(t *testing.T) {
	const Size = 16

	heap := S.New()
	insertele(heap, slice(0, Size))
	n := heap.Trees()
	if n != 1 {
		t.Errorf("something went wrong")
	}
	for i := 0; i < Size; i++ {
		pval := heap.Del_Min()
		if pval != i {
			t.Errorf("expected %d, got %d", i, pval)
		}
	}
}
func Test_Del_Min(t *testing.T) {
	heap := S.New()
	d := heap.Del_Min()
	if d != -1 {
		t.Error("expected", -1, "got", d)
	}
}

func Test_Del_Min1(t *testing.T) {
	heap := S.New()
	size := 100000
	for i := 0; i < size; i++ {
		n := size - 1 - i
		//fmt.Println("insert", n)
		heap.Insert(n)
	}
	for i := 0; i < size; i++ {
		n := heap.Del_Min()
		if n != i {
			t.Error("Expected", i, "Got", n)
		}
	}

	Values := []int{4, -90, 5, -248, 0, 23}
	DelOrder := []int{-248, -90, 0, 4, 5}
	for _, v := range Values {
		heap.Insert(v)
	}
	for _, v := range DelOrder {
		d := heap.Del_Min()
		if d != v {
			t.Error("expected", v, "Got", d)
		}
	}

}

func Test_Find_Min(t *testing.T) {
	heap := S.New()
	h := heap.Find_Min()
	if h != -1 {
		t.Error("Something went wrong")
	}
	Values := []int{4, -90, 5, -248, 0, 23}
	for _, v := range Values {
		heap.Insert(v)
	}
	d := heap.Find_Min()
	if d != -248 {
		t.Error("expected", -248, "Got", d)
	}

}

func Test_merge(t *testing.T) {
	Values := []int{9, 8, 7, 6, 5, 4, -3, -1000, 20000, -10000}
	Values1 := []int{3, 2, 1, 0, -1, -2, 300, 200}
	Values2 := []int{-10000, -1000, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 200, 300, 20000}
	size := len(Values) + len(Values1)
	h1 := S.New()
	h2 := S.New()
	for _, v := range Values {

		h1.Insert(v)
	}
	for _, v := range Values1 {

		h2.Insert(v)
	}
	h1.Merge(h2)
	h1.Print()
	fmt.Println("Size", h1.Size())
	if h1.Size() != size {
		t.Error("Something went wrong")
	}

	for i := 0; i < size; i++ {
		d := h1.Del_Min()
		fmt.Println("deleted:", d)
		if d != Values2[i] {
			t.Errorf("expected %d, got %d", Values2[i], d)
		}
	}
}

func slice(start int, end int) []int {

	slice := make([]int, 0, end-start)

	for i := start; i < end; i++ {
		slice = append(slice, i)
	}
	return slice
}

func insertele(bh *S.Binomial_Heap, values []int) {

	for _, elem := range values {
		bh.Insert(elem)
	}
}
