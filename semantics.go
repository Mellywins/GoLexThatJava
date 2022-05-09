package main

import (
	"fmt"
	"github.com/timtadh/lexmachine"
	"log"
)

type AbstractGlobalVariable struct {
	Value string
	Type  string
}
type AbstractLocalVariable struct {
	Value string
	Type  string
	Name  string
	State string // Has 2 states, either pristine (untouched, will throw an error if it reahes end of program without being used) or dirty(used in the code afterwards, this is the normal accepted behavior)
}
type FunctionSignature struct {
	Name       string
	ReturnType string
	ParamList  []AbstractLocalVariable
}

func NewFunctionSignature(name string, returnType string, paramList []AbstractLocalVariable) *FunctionSignature {
	return &FunctionSignature{Name: name, ReturnType: returnType, ParamList: paramList}
}

var CheckVariableHierarchyQueue []*Node
var GlobalVars = make(map[string]*AbstractGlobalVariable)
var LocalVars = make(map[string]*AbstractLocalVariable)
var FunctionDeclarations = make(map[string]*FunctionSignature)

func NewAbstractGlobalVariable(node *Node) *AbstractGlobalVariable {
	name, value, varType := destructureElementsFromVarDeclaration(node)
	constructedNode := &AbstractGlobalVariable{
		Value: value,
		Type:  varType,
	}
	GlobalVars[name] = constructedNode
	return constructedNode
}
func NewAbstractLocalVariable(val string, typ string, name string) *AbstractLocalVariable {
	return &AbstractLocalVariable{
		Value: val,
		Type:  typ,
		Name:  name,
		State: "Pristine",
	}
}
func destructureElementsFromVarDeclaration(node *Node) (name string, value string, varType string) {
	childrenCount := len(node.Children)
	tokenList := make([]*lexmachine.Token, childrenCount)
	for i, tok := range node.Children {
		tokenList[i] = tok.Token // [int x = 5 ; ]
	}
	if childrenCount > 3 { // uninitialized var
		return tokenList[1].Value.(string), tokenList[3].Value.(string), tokenList[0].Value.(string)
	} else {
		return tokenList[1].Value.(string), "", tokenList[0].Value.(string)
	}
}
func (n *Node) AddContext(node *Node) {
	name, value, varType := destructureElementsFromVarDeclaration(node)
	constructedNode := &AbstractLocalVariable{
		Value: value,
		Type:  varType,
	}
	n.Context[name] = constructedNode
}

// TODO: Try to find a way to backtrack to the nearest parent context.
func QueueElementForChecking(node *Node) {
	CheckVariableHierarchyQueue = append(CheckVariableHierarchyQueue, node)
}

// FindHealthyOccurenceInParentContexts Function that recursively checks for the variable existance within the parent scope until it reaches the parent = nil
func FindHealthyOccurenceInParentContexts(node *Node, name string) {

	if node.Parent == nil {
		if _, ok := node.Context[name]; !ok {
			log.Fatal(fmt.Errorf("variable %s is not defined previously", name))
			return
		} else {
			return
		}
	} else { // Incase parent exists, recurse and keep looking for occurence.
		FindHealthyOccurenceInParentContexts(node.Parent, name)
	}
}
