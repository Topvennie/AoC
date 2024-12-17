// Package gtree contains a generic tree implementation
package gtree

import (
	"fmt"
	"iter"
	"strings"
)

// Tree represents a generic tree
type Tree[T any] struct {
	Root   *Node[T]
	Size   int
	leaves []*Node[T]
}

// Node represents a node inside a generic tree
type Node[T any] struct {
	Value    T
	Children []*Node[T]
	tree     *Tree[T]
}

// New creates a new tree
func New[T any](value T) *Tree[T] {
	tree := &Tree[T]{
		Root:   nil,
		Size:   1,
		leaves: []*Node[T]{},
	}

	tree.Root = tree.NewNode(value)
	tree.leaves = append(tree.leaves, tree.Root)

	return tree
}

// NewNode creates a new Node
func (t *Tree[T]) NewNode(value T) *Node[T] {
	return &Node[T]{
		Value:    value,
		Children: []*Node[T]{},
		tree:     t,
	}
}

// Leaves returns an iterator over all the leafs of the tree
func (t *Tree[T]) Leaves() iter.Seq[*Node[T]] {
	if len(t.leaves) == 0 {
		return nil
	}

	idx := 0
	return func(yield func(*Node[T]) bool) {
		for yield(t.leaves[idx]) {
			idx++
			if idx >= len(t.leaves) {
				break
			}
		}
	}
}

func (t *Tree[T]) Print() string {
	var tree strings.Builder
	tree.WriteString(fmt.Sprintf("Size: %d\n", t.Size))
	tree.WriteString("Children:\n")

	children := []*Node[T]{t.Root}
	i := 0
	for {
		// Write node line
		var nodes strings.Builder
		nodes.WriteString(fmt.Sprintf("%.2d: ", i))

		newChildren := []*Node[T]{}

		for _, child := range children {
			newChildren = append(newChildren, child.Children...)
			nodes.WriteString(fmt.Sprintf("%+v |", child.Value))
		}

		tree.WriteString(fmt.Sprintf("%s\n", nodes.String()))
		if len(newChildren) == 0 {
			break
		}

		children = newChildren
	}

	return tree.String()
}

// AddChild adds a new child node to the node
func (n *Node[T]) AddChild(value T) *Node[T] {
	newNode := n.tree.NewNode(value)
	n.Children = append(n.Children, newNode)
	n.tree.Size++

	// Add to tree leaves
	idx := -1 // Index of this node
	for i, leaf := range n.tree.leaves {
		if leaf == n {
			idx = i
			break
		}
	}

	if idx != -1 {
		// Node was a leaf, not anymore
		n.tree.leaves[idx] = newNode
	} else {
		n.tree.leaves = append(n.tree.leaves, newNode)
	}

	return newNode
}
