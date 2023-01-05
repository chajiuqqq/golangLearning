/*
Map function with interfaces
Use the answer from exercise Q12, but now make it generic using interfaces. Make it at least work for ints and strings.
*/
package main

import "fmt"

type e interface{}

func Map(f func(e) e, x []e) []e {
	s := make([]e, len(x))
	for k, v := range x {
		s[k] = f(v)
	}
	return s
}

func main() {
	mult2 := func(a e) e {
		switch a.(type) {
		case int:
			return a.(int) * 2
		case string:
			return a.(string) + a.(string)
		}
		return a
	}
	a := []e{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", Map(mult2, a))
	b := []e{"ab", "bc", "cd", "de"}
	fmt.Printf("%v\n", Map(mult2, b))
}
