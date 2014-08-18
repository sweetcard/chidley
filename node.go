package main

import (
//"log"
)

type Node struct {
	name         string
	space        string
	spaceTag     string
	parent       *Node
	parents      []*Node
	children     map[string]*Node
	childCount   map[string]int
	repeats      bool
	nodeTypeInfo *NodeTypeInfo
	hasCharData  bool
	tempCharData string
}

type NodeVisitor interface {
	Visit(n *Node) bool
	AlreadyVisited(n *Node) bool
	SetAlreadyVisited(n *Node)
}

func (n *Node) initialize(name string, space string, spaceTag string, parent *Node) {
	n.parent = parent
	n.parents = make([]*Node, 0, 0)
	n.pushParent(parent)
	n.name = name
	n.space = space
	n.spaceTag = spaceTag
	n.children = make(map[string]*Node)
	n.childCount = make(map[string]int)
	n.nodeTypeInfo = new(NodeTypeInfo)
	n.nodeTypeInfo.initialize()
	n.hasCharData = false
}

func (n *Node) makeName() string {
	spaceTag := ""
	if n.spaceTag != "" {
		spaceTag = "_" + n.spaceTag
	}
	return capitalizeFirstLetter(cleanName(n.name)) + spaceTag
}

func (n *Node) makeType(prefix string, suffix string) string {
	return makeTypeGeneric(n.name, n.spaceTag, prefix, suffix)
}

func (n *Node) peekParent() *Node {
	if len(n.parents) == 0 {
		return nil
	}
	a := n.parents
	return a[len(a)-1]
}

func (n *Node) pushParent(parent *Node) {
	n.parents = append(n.parents, parent)
}

func (n *Node) popParent() *Node {
	if len(n.parents) == 0 {
		return nil
	}
	var poppedNode *Node
	a := n.parents
	poppedNode, n.parents = a[len(a)-1], a[:len(a)-1]
	return poppedNode
}

func makeTypeGeneric(name string, space string, prefix string, suffix string) string {
	spaceTag := ""
	if space != "" {
		spaceTag = space + "_"
	}
	if prefix != "" {
		prefix = prefix + "_"
	}

	if suffix != "" {
		suffix = "_" + suffix
	}
	return prefix + spaceTag + cleanName(name) + suffix

}