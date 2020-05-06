package tree

import "fmt"

type Node struct {
	Id    int
	Child []*Node
}

type Tree struct {
	Root *Node
}

// NewNode create new node and return pointer to it
// if child is passed as nil, error will be returned
func NewNode(id int, child []*Node) (*Node, error) {
	if child == nil {
		return nil, fmt.Errorf("nil input is passed for child")
	}
	return &Node{Id: id, Child: child}, nil
}

// PrintNode print node id value
func (n *Node) PrintNode() {
	fmt.Println(n.Id)
}

// AddChildren add children to given node
func (n *Node) AddChildren(children ...*Node) {
	n.Child = append(n.Child, children...)
}

// CreateTree create new tree and set root as input node
func CreateTree(root *Node) *Tree {
	return &Tree{Root: root}
}

// PrintPreOrder print all the node of tree in preorder
func (t *Tree) PrintPreOrder() {
	printPreorder(t.Root)
}

func printPreorder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Id)
	if len(root.Child) > 0 {
		for _, v := range root.Child {
			printPreorder(v)
		}
	}
}

// FindNode find node in the tree for the given id
// if not found, returns nil
func (t *Tree) FindNode(id int) *Node {
	return findNode(t.Root, id)
}

func findNode(root *Node, id int) *Node {

	if root == nil || root.Id == id {
		return root
	}

	if len(root.Child) > 0 {
		for _, v := range root.Child {
			n := findNode(v, id)
			if n != nil {
				return n
			}
		}
	}

	return nil
}

// FindShorestPath find shortest path of the node in the tree and
// returns it as []int, if not found, returns nil
func (t *Tree) FindShortestPath(id int) []int {
	path := []int{}
	found := findShortestPath(id, t.Root, &path)
	if found {
		return path
	}
	return nil
}

func findShortestPath(id int, root *Node, path *([]int)) bool {
	*path = append(*path, root.Id)
	if root.Id == id {
		return true
	}

	// look into all the children
	foundInChild := false
	for _, v := range root.Child {
		found := findShortestPath(id, v, path)
		if found {
			foundInChild = true
			break
		}
	}
	if foundInChild {
		return true
	} else {
		*path = (*path)[:len(*path)-1]
	}

	return false
}
