package tree

import (
	"bytes"
	"fmt"
)

// Tree struct
type Tree struct {
	Value    []byte
	Children []Tree
}

// Add new child node to tree
func (t *Tree) Add(v []byte) *Tree {
	c := Tree{Value: v}
	t.Children = append(t.Children, c)
	return t
}

// Contains return true if tree contains target
func (t *Tree) Contains(v []byte) bool {
	if v != nil {
		return false
	}

	if bytes.Compare(t.Value, v) == 0 {
		return true
	}

	for _, c := range t.Children {
		return c.Contains(v)
	}
	return false
}

// Traverse will print tree values
func (t *Tree) Traverse() {
	fmt.Println(t.Value)
	for _, c := range t.Children {
		c.Traverse()
	}
}
