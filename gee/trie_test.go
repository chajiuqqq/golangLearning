package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	n := NewNode()
	add(n, "/a/:b/*path")
	s := printNodes(n, 0)
	if s != "1:a 2::b 3:*path " {
		t.Fatal("error, s:", s)
	}

}

func printNodes(n *node, i int) (s string) {
	if n.part != "" {
		s += fmt.Sprintf("%d:%s ", i, n.part)
	}
	for _, node := range n.children {
		s += printNodes(node, i+1)
	}
	return
}

func TestPatternSplit(t *testing.T) {
	pattern := "/a/b/c"
	if !reflect.DeepEqual(patternSplit(pattern), []string{"a", "b", "c"}) {
		t.Fatal("not equal:", patternSplit(pattern))
	}
	if !reflect.DeepEqual(patternSplit("/"), []string{""}) {
		t.Fatal("not equal:", patternSplit("/"))
	}
}

func TestSearch(t *testing.T) {
	n := &node{children: make([]*node, 0)}
	add(n, "/:lang/doc")
	add(n, "/:lang/tutorial")
	add(n, "/:lang/help")
	add(n, "/about")
	add(n, "/p/*file")
	if searchPattern(n, "/en/doc") != "/:lang/doc" {
		t.Fatal("fail: /:lang/doc")
	}

	if searchPattern(n, "/en/tutorial") != "/:lang/tutorial" {
		t.Fatal("fail: /:lang/tutorial")
	}

	if searchPattern(n, "/en/help") != "/:lang/help" {
		t.Fatal("fail: /:lang/help")
	}

	if searchPattern(n, "/about") != "/about" {
		t.Fatal("fail: /about")
	}

	if searchPattern(n, "/p/asd/ds/file") != "/p/*file" {
		t.Fatal("fail: /p/*file")
	}

}

func add(n *node, s string) {
	n.insert(s, patternSplit(s))
}

func searchPattern(n *node, s string) string {
	if node := n.search(patternSplit(s), 0); node != nil {
		return node.pattern
	}
	return "nil"
}
