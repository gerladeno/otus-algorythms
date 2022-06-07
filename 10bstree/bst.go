package _0bstree

import (
	"fmt"
	"strconv"
)

type BST struct {
	Key    int
	Value  []interface{}
	Parent *BST
	Left   *BST
	Right  *BST
}

func (t *BST) Insert(key int, value interface{}) {
	switch {
	case t.Key == key:
		t.Value = append(t.Value, value)
	case key < t.Key && t.Left == nil:
		t.Left = &BST{
			Key:    key,
			Value:  []interface{}{value},
			Parent: t,
		}
	case key < t.Key && t.Left != nil:
		t.Left.Insert(key, value)
	case key > t.Key && t.Right == nil:
		t.Right = &BST{
			Key:    key,
			Value:  []interface{}{value},
			Parent: t,
		}
	case key > t.Key && t.Right != nil:
		t.Right.Insert(key, value)
	}
}

func (t *BST) Search(key int) *BST {
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

func (t *BST) Delete(key int) bool {
	target := t.Search(key)
	if target == nil {
		return false
	}
	switch {
	case target.Left == nil && target.Right == nil && target.Parent == nil:
		return false
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
		return false
	}
	return true
}

func (t *BST) linkLeft(target *BST) {
	t.Left = target
	target.Parent = t
}

func (t *BST) linkRight(target *BST) {
	t.Right = target
	target.Parent = t
}

func (t *BST) deepestRight() *BST {
	if t.Right == nil {
		return t
	}
	return t.Right.deepestRight()
}

func (t *BST) Draw() {
	drawer := [][]byte{[]byte(" " + strconv.Itoa(t.Key))}
	t.drawLevel(&drawer, 1)
	for _, str := range drawer {
		fmt.Println(string(str))
	}
}

func (t *BST) drawLevel(drawer *[][]byte, level int) {
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
