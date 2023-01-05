/*
Functions that return functions
*/
package main
import "fmt"
func plusTwo() func(int) int {
	f:=func(a int) int{
		return a+2
	}
	return f
}

func plusX(x int) func(int) int {
	f:=func(a int) int{
		return a+x
	}
	return f
}

func main(){
	s:=plusTwo()
	fmt.Printf("%v\n",s(3))
	s2:=plusX(10)
	fmt.Printf("%v\n",s2(3))
}