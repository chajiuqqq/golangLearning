/*
Create your own linked list implementation
*/
package main

import (
	"container/list"
	"fmt"
)

func q1() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(4)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
}

//q2

type any interface{}
type MyElem struct {
	Value          any
	Previous, Next *MyElem
}

type MyList struct {
	head, tail *MyElem
}

func (l *MyList) PushBack(v any) *MyList {
	n := &MyElem{Value: v}
	if l.head == nil {
		l.head = n
	} else {
		l.tail.Next = n
		n.Previous = l.tail
	}
	l.tail = n
	return l
}

func (l *MyList) Pop() *MyElem {
	n := l.tail
	if l.tail != l.head {
		l.tail = n.Previous
	} else {
		l.head = nil
		l.tail = nil
	}
	return n
}

func main() {
	l := &MyList{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(4)
	for e := l.head; e != nil; e = e.Next {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()

	for e := l.Pop(); e != nil; e = l.Pop() {
		fmt.Printf("%v\n", e.Value)
	}

}
