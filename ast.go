package main

import (
	"fmt"
	"github.com/timtadh/lexmachine"
	"strings"
)

type Node struct {
	Name     string
	Token    *lexmachine.Token
	Children []*Node
	parent   *Node
}

const (
	newLine      = "\n"
	emptySpace   = "    "
	middleItem   = "├── "
	continueItem = "│   "
	lastItem     = "└── "
)

func NewNode(name string, token *lexmachine.Token) *Node {
	return &Node{
		Name:  name,
		Token: token,
	}
}

// Add a child at the end of the children list to a node.
func (n *Node) AddKid(kid *Node) *Node {
	n.Children = append(n.Children, kid)
	if kid != nil {
		kid.parent = n
	}
	return n
}

// Add a child at the start of the children list to a Node
func (n *Node) PrependKid(kid *Node) *Node {
	kids := append(make([]*Node, 0, cap(n.Children)+1), kid)
	n.Children = append(kids, n.Children...)
	if kid != nil {
		kid.parent = n
	}
	return n
}

// String makes the AST human readable

func (n *Node) String() string {
	if n == nil {
		return ""
	}
	parts := make([]string, 0, len(n.Children))
	parts = append(parts, n.Name)
	if n.Token != nil && string(n.Token.Lexeme) != n.Name {
		parts = append(parts, fmt.Sprintf("%q", string(n.Token.Lexeme)))
	}
	for _, k := range n.Children {
		parts = append(parts, k.String())
	}
	if len(parts) > 1 {
		return fmt.Sprintf("%v", strings.Join(parts, " "))
	}
	return strings.Join(parts, " ")
	//p := printer{}
	//return p.Print(*n)
}

//
//type printer struct {
//}
//type PrintableNode interface {
//	Print(Node) string
//}
//
//func (n *Node) Text() string {
//	if n.Token == nil {
//		return n.Name
//	}
//	return string(n.Token.Lexeme)
//}
//func (n *Node) Items() []*Node {
//	return n.Children
//}
//
//func (p *printer) printText(text string, spaces []bool, last bool) string {
//	var result string
//	for _, space := range spaces {
//		if space {
//			result += emptySpace
//		} else {
//			result += continueItem
//		}
//	}
//
//	indicator := middleItem
//	if last {
//		indicator = lastItem
//	}
//
//	var out string
//	lines := strings.Split(text, "\n")
//	for i := range lines {
//		text := lines[i]
//		if i == 0 {
//			out += result + indicator + text + newLine
//			continue
//		}
//		if last {
//			indicator = emptySpace
//		} else {
//			indicator = continueItem
//		}
//		out += result + indicator + text + newLine
//	}
//
//	return out
//}
//
//func (p *printer) printItems(t []*Node, spaces []bool) string {
//	var result string
//	for i, f := range t {
//		last := i == len(t)-1
//		result += p.printText(f.Text(), spaces, last)
//		if len(f.Items()) > 0 {
//			spacesChild := append(spaces, last)
//			result += p.printItems(f.Items(), spacesChild)
//		}
//	}
//	return result
//}
//func (p *printer) Print(t Node) string {
//	return t.Text() + newLine + p.printItems(t.Items(), []bool{})
//}
