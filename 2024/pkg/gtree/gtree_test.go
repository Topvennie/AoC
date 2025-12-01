package gtree_test

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/Topvennie/AoC/pkg/gtree"
)

func TestNew(t *testing.T) {
	// Test int iTree
	iTree := gtree.New(0)

	if iTree.Size != 1 {
		t.Errorf("Incorrect size, expected 1 got: %d", iTree.Size)
	}

	if iTree.Root.Value != 0 {
		t.Errorf("Incorrect root value, expected 0 got: %d", iTree.Root.Value)
	}

	// Test string sTree
	sTree := gtree.New("tree")
	if sTree.Root.Value != "tree" {
		t.Errorf("Incorrect root value, exepected 'tree' got: %s", sTree.Root.Value)
	}
}

func TestNewNode(t *testing.T) {
	tree := gtree.New(-1)
	node := tree.NewNode(0)

	if node.Value != 0 {
		t.Errorf("Incorrect node value, expected 0 got: %d", node.Value)
	}
	if len(node.Children) != 0 {
		t.Errorf("Incorrect children size, expected 0 got: %d", len(node.Children))
	}

	for range 1000 {
		value := -500 + rand.Intn(1000) // Random number between -500 and 500
		node = tree.NewNode(value)

		if node.Value != value {
			t.Errorf("Incorrect node value, expected %d got: %d", value, node.Value)
		}
	}
}

func TestLeaves(t *testing.T) {
	tree := gtree.New(-1)

	// Test iterator with only root
	i := 0
	for l := range tree.Leaves() {
		i++
		if l != tree.Root {
			t.Error("Leaves iterator is incorrect for a tree with only a root")
		}
	}
	if i != 1 {
		t.Errorf("Leaves iterator size is incorrect for a tree with only a root, expected 1 got: %d", i)
	}

	// 0:        R
	// 1:    1       2
	// 2:  3   4   5   6
	// 3: 7 8 9

	// 1
	one := tree.Root.AddChild(1)
	two := tree.Root.AddChild(2)

	// 2
	three := one.AddChild(3)
	four := one.AddChild(4)
	five := two.AddChild(5)
	six := two.AddChild(6)

	// 3
	seven := three.AddChild(7)
	eight := three.AddChild(8)
	nine := four.AddChild(9)

	leaves := []*gtree.Node[int]{seven, eight, nine, five, six}
	for l := range tree.Leaves() {
		deleted := false
		for i, leaf := range leaves {
			if l == leaf {
				leaves = append(leaves[:i], leaves[i+1:]...)
				deleted = true
				break
			}
		}

		if !deleted {
			t.Errorf("Leaves iterator contained an incorect node with value: %d", l.Value)
		}
	}

	if len(leaves) != 0 {
		t.Errorf("Leaves iterator did not contain the nodes: %+v", leaves)
	}
}

func TestAddChild(t *testing.T) {
	tree := gtree.New(-1)

	// 0:          R
	// 1:   1      2      3
	// 2: 4 5 6  7

	// 1
	one := tree.Root.AddChild(1)
	two := tree.Root.AddChild(2)
	three := tree.Root.AddChild(3)

	// 2
	four := one.AddChild(4)
	five := one.AddChild(5)
	six := one.AddChild(6)
	seven := two.AddChild(7)

	if tree.Size != 8 {
		t.Errorf("Incorrect tree size, expected 8 got: %d", tree.Size)
	}

	rootChildren := []*gtree.Node[int]{one, two, three}
	if !slices.Equal(rootChildren, tree.Root.Children) {
		t.Errorf("Root children incorrect, expected %+v got: %+v", rootChildren, tree.Root.Children)
	}

	oneChildren := []*gtree.Node[int]{four, five, six}
	if !slices.Equal(oneChildren, one.Children) {
		t.Errorf("One children incorrect, expected %+v got: %+v", oneChildren, one.Children)
	}

	twoChildren := []*gtree.Node[int]{seven}
	if !slices.Equal(twoChildren, two.Children) {
		t.Errorf("Two children incorrect, expected %+v got: %+v", twoChildren, two.Children)
	}

	emptyChildren := []*gtree.Node[int]{}
	if !(slices.Equal(emptyChildren, three.Children) &&
		slices.Equal(emptyChildren, four.Children) &&
		slices.Equal(emptyChildren, five.Children) &&
		slices.Equal(emptyChildren, six.Children) &&
		slices.Equal(emptyChildren, seven.Children)) {
		t.Errorf("Empty children incorrect")
	}
}
