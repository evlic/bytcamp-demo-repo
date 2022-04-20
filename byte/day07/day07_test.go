package day07

import (
	"fmt"
	"testing"
)

func TestDay07(t *testing.T) {
	testArr, K := []int{100, 200, 34, 53, 452, 3425, 5345, 23, 43, 534, 2, 123}, 3
	linked := NewLinedListFromIntS(testArr)
	h := NewNodeHeap(K)

	for p := linked.Head.Next; p != nil ; p = p.Next{
		h.NodePush(p)
	}

	// for h.Len() > 0 {
	// 	fmt.Printf("\t%v\n", heap.Pop(h).(*Node).Val)
	// }
	fmt.Println(h)
}
