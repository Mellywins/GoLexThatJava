package syntax

import "github.com/timtadh/lexmachine"

type Node struct {
	Name     string
	Token    *lexmachine.Token
	Children []*Node
}

func newNode(name string, token *lexmachine.Token) *Node {
	return &Node{
		Name:  name,
		Token: token,
	}
}

// Add a child at the end of the children list to a node.
func (n *Node) AddKid(kid *Node) *Node {
	n.Children = append(n.Children, n)
	return n
}

// Add a child at the start of the children list to a Node
func (n *Node) PrependKid(kid *Node) *Node {
	kids := append(make([]*Node, 0, cap(n.Children)+1), kid)
	n.Children = append(kids, n.Children...)
	return n
}

// String makes the AST human readable

func (n *Node) String() string {
	parts := make([]string, 0, len(n.Children))
	parts = append(parts, n.Name)
}
