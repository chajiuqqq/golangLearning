/*
Write a function that finds the maximum value in an int slice ([]int).
Write a function that finds the minimum value in an int slice ([]int).
*/
package main

import "fmt"
func maximum(numbers []int) int{
	if len(numbers)==0{
		return 0
	}
	m := numbers[0]
	for _,num := range numbers{
		if num>m{
			m=num
		}
	}
	return m
}
func minimum(numbers []int) int{
	
	if len(numbers)==0{
		return 0
	}
	m := numbers[0]
	for _,num := range numbers{
		if num<m{
			m=num
		}
	}
	return m
}

func main(){
	s := []int{1,2,3,4,5}
	fmt.Printf("slices:%v,max:%d,min:%d",s,maximum(s),minimum(s))
}