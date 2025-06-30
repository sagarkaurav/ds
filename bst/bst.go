package bst

import (
	"fmt"

	"github.com/sagarkaurav/ds/queue"
	"golang.org/x/exp/constraints"
)

type node[T constraints.Ordered] struct {
	val   T
	left  *node[T]
	right *node[T]
}

type BST[T constraints.Ordered] struct {
	root *node[T]
}

func New[T constraints.Ordered]() *BST[T] {
	return &BST[T]{root: nil}
}

func insert[T constraints.Ordered](n *node[T], val T) *node[T] {
	newNode := &node[T]{val: val}
	if n == nil {
		return newNode
	}
	if n.val > val {
		n.left = insert(n.left, val)
	}
	if n.val < val {
		n.right = insert(n.right, val)
	}
	return n
}

func (bst *BST[T]) Insert(val T) {
	bst.root = insert(bst.root, val)
}

func (bst *BST[T]) Delete(val T) {
	bst.root = delete(bst.root, val)
}

func delete[T constraints.Ordered](n *node[T], val T) *node[T] {
	if n == nil {
		return nil
	}
	if n.val > val {
		n.right = delete(n.right, val)
	}
	if n.val < val {
		n.left = delete(n.left, val)
	}
	if n.val == val {
		leftTree := n.left
		rightTree := n.right
		if leftTree == nil {
			return rightTree
		}
		if rightTree == nil {
			return leftTree
		}
		currentNode := leftTree
		for currentNode.right != nil {
			currentNode = currentNode.right
		}
		if currentNode != nil {
			leftTree = delete(leftTree, currentNode.val)
		}
		currentNode.left = leftTree
		currentNode.right = rightTree
		return currentNode
	}
	return n
}

func (bst *BST[T]) LevelOrderPrint() {
	queue := queue.New[*node[T]]()
	queue.Enqueue(bst.root)
	for queue.Size() != 0 {
		node := queue.Dequeue()
		if node.left != nil {
			queue.Enqueue(node.left)
		}
		if node.right != nil {
			queue.Enqueue(node.right)
		}
		fmt.Println(node.val)
	}
}
