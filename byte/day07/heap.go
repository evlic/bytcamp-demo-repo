package day07

import (
	"container/heap"
	"fmt"
)

const (
	NodeNilErr = "Node 不可以为空"
)

func (h NodeHeap) Len() int { return len(h.heap) }

// Less 当 h[i] < h[j] return true 不执行交换，
// Less 方法可以理解为 isNotSwap，true 不交换，如果此时 i < j 则为升序
func (h NodeHeap) Less(i, j int) bool { return h.heap[i].Val < h.heap[j].Val }
func (h NodeHeap) Swap(i, j int)      { h.heap[i], h.heap[j] = h.heap[j], h.heap[i] }
func (h *NodeHeap) Pop() any {
	n := h.Len() - 1
	lastNode := h.heap[n]
	h.heap = h.heap[:n]
	return lastNode
}
func (h *NodeHeap) Push(x any) {
	var node *Node
	if node = x.(*Node); node == nil {
		panic(NodeNilErr)
	}
	h.heap = append(h.heap, node)
}

// NodeHeap 链表节点堆，题目求取 TopK 大则使用小顶堆实现。
type NodeHeap struct {
	heap []*Node
	// 不对外提供 set
	kLimit int
}

func NewNodeHeap(K int) *NodeHeap {
	return &NodeHeap{make([]*Node, 0, K), K}
}

// NodePush 封装 heap.Push 平衡堆的调整
// 并嵌入操作逻辑，判断是否达到上限 TopK，决定调用 push 还是替换堆顶后 fix
func (h *NodeHeap) NodePush(node *Node) {
	if h.Len() < h.kLimit {
		heap.Push(h, node)
	} else if h.Top() < node.Val {
		// 如果堆顶元素比将插入的元素还要小，则对其替换并 fix
		h.heap[0] = node
		heap.Fix(h, 0)
	}
}

func (h NodeHeap) Top() int {
	return h.heap[0].Val
}

func (h *NodeHeap) String() (res string) {
	res += "NodeHeap{"
	m := "\n\t「%v」%v"
	for i, node := range h.heap {
		res += fmt.Sprintf(m, i, node.Val)
	}
	res += "\n}"
	return
}
