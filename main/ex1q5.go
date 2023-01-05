package main

import "fmt"
func cal_avg(s []float64)  {
	var sum float64
	for _,t := range s{
		sum += t
	}
	//len()返回int，还得转为float64才能和sum运算
	fmt.Printf("sum:%f,avg:%f",sum,sum/float64(len(s)))
}
func main()  {
	cal_avg([]float64{123.2,45.23,1002.3})
}