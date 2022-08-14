package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func main() {
	myT := &TreeNode{val: 8}
	myT.insert(4)
	myT.insert(2)
	myT.insert(3)
	myT.insert(10)
	myT.insert(6)
	myT.insert(7)
	myT.printOrder()
	fmt.Println()

	fmt.Println(myT.min())
	fmt.Println(myT.max())
	myT.delete(3)
	fmt.Println()
	myT.printOrder()
}

// Delete function
func (t *TreeNode) delete(element int) {
	t.remove(element)
}

func (t *TreeNode) remove(element int) *TreeNode {
	if t == nil {
		return nil
	}

	if element < t.val {
		t.left = t.left.remove(element)
		return t
	}
	if element > t.val {
		t.right = t.right.remove(element)
		return t
	}

	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}

	if t.left == nil {
		t = t.right
		return t
	}
	if t.right == nil {
		t = t.left
		return t
	}

	minRight := t.right
	for {
		if minRight != nil && minRight.left != nil {
			minRight = minRight.left
		} else {
			break
		}
	}

	t.val = minRight.val
	t.right = t.right.remove(t.val)
	return t
}

func (t *TreeNode) printOrder() {
	if t == nil {
		return
	}
	t.left.printOrder()
	fmt.Println(t.val)
	t.right.printOrder()
}

func (t *TreeNode) insert(element int) {
	if t == nil {
		return
	}
	if element < t.val {
		if t.left == nil {
			t.left = &TreeNode{val: element}
		} else {
			t.left.insert(element)
		}
	} else {
		if t.right == nil {
			t.right = &TreeNode{val: element}
		} else {
			t.right.insert(element)
		}
	}
}

// Minimum value function
func (t *TreeNode) min() int {
	if t.left == nil {
		return t.val
	}
	return t.left.min()
}

// Maximum value function
func (t *TreeNode) max() int {
	if t.right == nil {
		return t.val
	}
	return t.right.max()
}
