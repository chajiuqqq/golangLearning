package main

import (
	"fmt"
	"unicode/utf8"
)
func s1()  {
	for i:=0;i<100;i=i+1{
		for j:=0;j<=i;j=j+1{
			fmt.Printf("A")
		}
		fmt.Printf("\n")
	}
}

func s2()  {
	s := "asSASA ddd dsjkdsjs dk中国"
	b := []byte(s)
	fmt.Printf("strlen:%d,byte lens:%d, rune lens:%d\n",len(s),len(b),utf8.RuneCount(b))
	//strlen:28,byte lens:28, rune lens:24
	// len(s) 获得byte数
}

func s3()  {
	s := "asSASA ddd dsjkdsjs dk中国"
	r := []rune(s)
	r[4] = 'a'
	r[5] = 'b'
	r[6] = 'c'

	fmt.Printf(string(r))
}

//用copy的方法替换字符
func s3_1()  {
	s := "asSASA ddd dsjkdsjs dk中国"
	r := []rune(s)
	copy(r[4:4+3],[]rune("abc"))

	fmt.Printf("before:%s\n",s)
	fmt.Printf("after:%s\n",string(r))
}

func s4()  {
	s := "foobar"
	r := []rune(s)
	//获取rune个数，可以转为[]rune后用len()
	for i,j:=0,len(r)-1;i<j;i,j=i+1,j-1{
		//可以直接交换
		r[i],r[j]=r[j],r[i]
	}

	fmt.Printf("before:%s\n",s)
	fmt.Printf("after:%s\n",string(r))
}


func main()  {
s4()
}