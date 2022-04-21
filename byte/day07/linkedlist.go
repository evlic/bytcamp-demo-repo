package day07

const (
	HeadT = 0xefffffff
)

type LinkedList struct {
	Head, Tail *Node
	Len        int
}

func (l *LinkedList) Push(node *Node) {
	if node == nil {
		panic(NodeNilErr)
	}
	l.Tail.Next = node
	l.Tail = node
	l.Len++

}

func NewLinedListFromIntS(arr []int) *LinkedList {
	h := &Node{HeadT, nil}
	product := &LinkedList{h, h, 0}
	for _, i := range arr {
		product.Push(&Node{i, nil})
	}
	return product
}

// Node 单链表节点
type Node struct {
	Val  int
	Next *Node
}
