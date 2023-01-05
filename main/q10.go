/*Write a function that takes a variable number of ints and prints each integer on a separate line*/
package main

import "fmt"
func printInt(ints ...int){
	for _,item := range ints{
		fmt.Printf("%d\n",item)
	}
	
}
func main(){
	printInt(1,2,3,4,4,5,6)
}