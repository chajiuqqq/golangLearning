/*
create a program that shows the maximum number and that works for both integers and floats.
Try to make your program as generic as possible,
*/
package main

import "fmt"

type any interface{}

func less(a, b any) bool {
	switch a.(type) {
	case int:
		if _, ok := a.(int); ok {
			return a.(int) < b.(int)
		}
	case float32:
		if _, ok := a.(float32); ok {
			return a.(float32) < b.(float32)
		}
	}
	return false

}

func maximum(numbers []any) any {
	a := numbers[0]
	for _, num := range numbers {
		if less(a, num) {
			a = num
		}
	}
	return a
}

func main() {
	s := []any{1, 2, 3, 4, 5}
	fmt.Printf("slices:%v,max:%d\n", s, maximum(s))

	s1 := []any{1.2, 2.3, 1.3, 6.4, 5.5}
	fmt.Printf("slices:%v,max:%f\n", s1, maximum(s1))

}
