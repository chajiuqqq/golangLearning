/*
Map function A map()-function is a function that takes a function and a list. 
The function is applied to each member in the list and a new list containing these calculated values is returned.

Write a simple map()-function in Go.
*/
package main

import "fmt"
func Map(f func(int) int, numbers []int) []int{
	s := make([]int,len(numbers))
	for i,num := range numbers{
		s[i]=f(num)
	}
	return s
}

func main(){
	f := func(a int) int{
		return a+1
	}
	fmt.Printf("%v",Map(f,[]int{1,2,3,4,5}))
}