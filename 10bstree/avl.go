package _0bstree

import (
	"fmt"
	"strconv"
)

type AVLTree struct {
	Key    int
	Value  []interface{}
	Parent *AVLTree
	Left   *AVLTree
	Right  *AVLTree
	Height int
}

func (t *AVLTree) Insert(key int, value interface{}) *AVLTree {
	switch {
	case t.Key == key:
		t.Value = append(t.Value, value)
	case key < t.Key && t.Left == nil:
		t.Left = &AVLTree{
			Key:    key,
			Value:  []interface{}{value},
			Parent: t,
			Height: 1,
		}
	case key < t.Key && t.Left != nil:
		t.Left.Insert(key, value)
	case key > t.Key && t.Right == nil:
		t.Right = &AVLTree{
			Key:    key,
			Value:  []interface{}{value},
			Parent: t,
			Height: 1,
		}
	case key > t.Key && t.Right != nil:
		t.Right.Insert(key, value)
	}
	return t.rebalanceTree()
}

func (t *AVLTree) Search(key int) *AVLTree {
	switch {
	case t.Key == key:
		return t
	case t.Key > key && t.Left != nil:
		return t.Left.Search(key)
	case t.Key < key && t.Right != nil:
		return t.Right.Search(key)
	}
	return nil
}

func (t *AVLTree) Delete(key int) *AVLTree {
	target := t.Search(key)
	if target == nil {
		return t
	}
	switch {
	case target.Left == nil && target.Right == nil && target.Parent == nil:
		return t
	case target.Left == nil && target.Right == nil && target.Parent.Left == target:
		target.Parent.Left = nil
	case target.Left == nil && target.Right == nil && target.Parent.Right == target:
		target.Parent.Right = nil
	case target.Right == nil && target.Left != nil && target.Parent != nil && target.Parent.Left == target:
		target.Parent.linkLeft(target.Left)
	case target.Right == nil && target.Left != nil && target.Parent != nil && target.Parent.Right == target:
		target.Parent.linkRight(target.Left)
	case target.Left == nil && target.Right != nil && target.Parent != nil && target.Parent.Left == target:
		target.Parent.linkLeft(target.Right)
	case target.Left == nil && target.Right != nil && target.Parent != nil && target.Parent.Right == target:
		target.Parent.linkRight(target.Right)
	case target.Left != nil && target.Right != nil && target.Parent == nil:
		deepestRight := target.Left.deepestRight()
		deepestRight.linkRight(target.Right)
		*t = *target.Left
	case target.Left != nil && target.Right != nil && target.Parent != nil:
		deepestRight := target.Left.deepestRight()
		deepestRight.linkRight(target.Right)
		if target.Parent.Left == target {
			target.Parent.linkLeft(target.Left)
		} else {
			target.Parent.linkRight(target.Left)
		}
	default:
		return t
	}
	return t.rebalanceTree()
}

func (t *AVLTree) linkLeft(target *AVLTree) {
	t.Left = target
	target.Parent = t
}

func (t *AVLTree) linkRight(target *AVLTree) {
	t.Right = target
	target.Parent = t
}

func (t *AVLTree) deepestRight() *AVLTree {
	if t.Right == nil {
		return t
	}
	return t.Right.deepestRight()
}

func (t *AVLTree) getHeight() int {
	if t == nil {
		return 0
	}
	return t.Height
}

func (t *AVLTree) recalculateHeight() {
	t.Height = 1 + max(t.Left.getHeight(), t.Right.getHeight())
}

func (t *AVLTree) rebalanceTree() *AVLTree {
	if t == nil {
		return t
	}
	t.recalculateHeight()
	balanceFactor := t.Left.getHeight() - t.Right.getHeight()
	if balanceFactor == -2 {
		if t.Right.Left.getHeight() > t.Right.Right.getHeight() {
			t.Right = t.Right.rotateRight()
		}
		return t.rotateLeft()
	} else if balanceFactor == 2 {
		if t.Left.Right.getHeight() > t.Left.Left.getHeight() {
			t.Left = t.Left.rotateLeft()
		}
		return t.rotateRight()
	}
	return t
}

func (t *AVLTree) rotateLeft() *AVLTree {
	newRoot := t.Right
	t.Right = newRoot.Left
	newRoot.Left = t
	t.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (t *AVLTree) rotateRight() *AVLTree {
	newRoot := t.Left
	t.Left = newRoot.Right
	newRoot.Right = t
	t.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (t *AVLTree) Draw() {
	drawer := [][]byte{[]byte(" " + strconv.Itoa(t.Key))}
	t.drawLevel(&drawer, 1)
	for _, str := range drawer {
		fmt.Println(string(str))
	}
}

func (t *AVLTree) drawLevel(drawer *[][]byte, level int) {
	if drawer == nil || len(*drawer) < level {
		return
	}
	if len(*drawer) < level+1 {
		*drawer = append(*drawer, []byte{})
	}
	if t.Left == nil {
		(*drawer)[level] = append((*drawer)[level], []byte(" X")...)
	} else {
		(*drawer)[level] = append((*drawer)[level], []byte(" "+strconv.Itoa(t.Left.Key))...)
		t.Left.drawLevel(drawer, level+1)
	}
	if t.Right == nil {
		(*drawer)[level] = append((*drawer)[level], []byte(" X")...)
	} else {
		(*drawer)[level] = append((*drawer)[level], []byte(" "+strconv.Itoa(t.Right.Key))...)
		t.Right.drawLevel(drawer, level+1)
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
