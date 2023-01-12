package gee

import (
	"log"
	"strings"
)

type node struct {
	pattern  string
	part     string
	children []*node
	wild     bool
}

func NewNode() *node {
	return &node{}
}

func patternSplit(pattern string) []string {
	return strings.Split(pattern, "/")[1:]
}

// 插入时寻找n的第一个符合part的child node（完全匹配）
func (n *node) searchChild(part string) *node {

	for _, child := range n.children {
		if child.part == part {
			return child
		}
	}
	return nil

}

// 搜索时寻找n所有符合part的children nodes（可以通配符）
func (n *node) searchChildren(part string) []*node {

	nodes := make([]*node, 0)

	for _, child := range n.children {
		if child.part == part || child.wild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string) {
	foundChild := n
	for _, part := range parts {
		nextNode := foundChild.searchChild(part)
		if nextNode == nil {
			nextNode = &node{pattern: pattern, part: part, wild: part != "" && (part[0] == ':' || part[0] == '*')}
			foundChild.children = append(foundChild.children, nextNode)
		}
		foundChild = nextNode
		if foundChild == nil {
			log.Fatal("foundChild == nil,part:", part)
		}

	}
}

func (n *node) search(parts []string, depth int) *node {
	//如果待匹配part没了，或者node是*
	if len(parts) == depth || strings.HasPrefix(n.part, "*") {
		//确保是叶子节点
		if len(n.children) != 0 {
			return nil
		}
		return n
	}
	nodes := n.searchChildren(parts[depth])
	for _, node := range nodes {
		if r := node.search(parts, depth+1); r != nil {
			return r
		}
	}
	return nil
}
