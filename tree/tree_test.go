package tree

import (
	"errors"
	"testing"
)

const (
	errorMsg = "%v : Expected %v, Got %v"
)

func TestNewNode(t *testing.T) {
	child1 := &Node{
		Id:    10,
		Child: []*Node{},
	}

	child2 := &Node{
		Id:    11,
		Child: []*Node{},
	}

	test := []struct {
		input  Node
		output error
	}{
		{
			input: Node{
				Id:    1,
				Child: nil,
			},
			output: errors.New("nil input is passed for child"),
		},
		{
			input: Node{
				Id:    1,
				Child: []*Node{},
			},
			output: nil,
		},
		{
			input: Node{
				Id:    2,
				Child: []*Node{child1},
			},
			output: nil,
		},
		{
			input: Node{
				Id:    3,
				Child: []*Node{child1, child2},
			},
			output: nil,
		},
	}

	for _, v := range test {
		out, err := NewNode(v.input.Id, v.input.Child)
		if v.input.Child == nil {
			if err.Error() != v.output.Error() {
				t.Errorf(errorMsg, "NewNode - error", v.output.Error(), err.Error())
				continue
			}
		} else if out == nil {
			t.Errorf(errorMsg, "NewNode - node", "!nil", "nil")
		}
	}
}

func Example_printNode() {
	n := &Node{
		Id:    1,
		Child: []*Node{},
	}
	n.PrintNode()

	// Output
	// 1
}

func TestAddChildren(t *testing.T) {
	parent := &Node{
		Id:    10,
		Child: []*Node{},
	}

	child2 := &Node{
		Id:    11,
		Child: []*Node{},
	}

	parent.AddChildren(child2)

	if len(parent.Child) == 1 {
		if parent.Child[0].Id != 11 {
			t.Errorf("expected %v,got %v for child id", 11, parent.Child[0].Id)
		}
	} else {
		t.Errorf("expected %v,got %v for child numbers", 1, len(parent.Child))
	}
}

func TestCreateTree(t *testing.T) {
	root := &Node{Id: 1}
	tree := CreateTree(root)
	if tree != nil {
		if tree.Root.Id != 1 {
			t.Errorf("expected %v, got %v for root id", 1, tree.Root.Id)
		}
	} else {
		t.Errorf("expected %v, got %v for root", "not nil", nil)
	}
}

func Example_printPreOrder() {
	node1 := &Node{Id: 1}
	node2 := &Node{Id: 2}
	node3 := &Node{Id: 3}
	node1.Child = append(node1.Child, node2, node3)

	tree := CreateTree(node1)
	tree.PrintPreOrder()
	// Output
	// 1
	// 2
	// 3
}

func TestFindNode(t *testing.T) {
	node1 := &Node{Id: 1}
	node2 := &Node{Id: 2}
	node3 := &Node{Id: 3}
	node1.Child = append(node1.Child, node2, node3)

	tree := CreateTree(node1)
	n := tree.FindNode(1)
	if n != nil {
		if n.Id != 1 {
			t.Errorf("expected %v, got %v for node id", 1, n.Id)
		}
	} else {
		t.Errorf("expected %v, got %v for node", "not nil", nil)
	}

}

func TestFindShotestPath(t *testing.T) {
	node1 := &Node{Id: 0}
	node2 := &Node{Id: 1}
	node3 := &Node{Id: 2}
	node1.Child = append(node1.Child, node2, node3)

	tree := CreateTree(node1)
	path := tree.FindShortestPath(1)

	if len(path) != 2 {
		t.Errorf("expected %v, got %v for path length", 2, len(path))
	}

	for i, v := range path {
		if i != v {
			t.Errorf("expected %v, got %v for path value", i+1, v)
		}
	}
}
