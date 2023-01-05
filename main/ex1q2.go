package main

import "fmt"

func s1(){
	fmt.Printf("s1:\n")
	for i:=0;i<10;i=i+1{
		fmt.Printf("%d\n",i)
	}
}

func s2(){
	
	fmt.Printf("s2:\n")
	i:=0
printLabel:
	if i<10{
		fmt.Printf("%d\n",i)
		i=i+1
		goto printLabel
	}	

}


func s3(){
	var a [10]int
	for i:=0;i<10;i=i+1{
		a[i]=i
	}
	fmt.Printf("s3:\n")
	fmt.Printf("%v\n",a)
	

}

func main(){
	s3()
}