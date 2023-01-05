/*
The Fibonacci sequence starts as follows: 1; 1; 2; 3; 5; 8; 13; : : : 
Or in mathematical terms: x1 = 1; x2 = 1; xn = xn 1 + xn 2 8n > 2. Write a function that takes an int value and gives that many terms of the Fibonacci sequence.
*/
package main

import "fmt"
func f(n int){
	s := make([]int,n)
	s[0],s[1]=1,1
	for i:=2;i<n;i++{
		s[i]=s[i-2]+s[i-1]
	}
	fmt.Printf("%v\n",s)
}
func main(){
	f(7)
}