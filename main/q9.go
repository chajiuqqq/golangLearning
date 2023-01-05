package main
import "fmt"
var loc=0
func push(s []int, v int) bool{
	if len(s)==loc{
		return false
	}
	s[loc]=v
	loc+=1
	return  true
}

func pop(s []int) ( val int,success bool){
	if loc==0{
		return
	}
	loc-=1
	val = s[loc]
	success = true
	return 
}


func main()  {
	s:=make([]int,5)

	push(s,1)
	push(s,2)
	push(s,3)
	push(s,4)
	push(s,5)
	fmt.Printf("s:%v\n",s)
	a,ok :=pop(s)
	if ok{
		fmt.Printf("pop:%v\ns: ",a)
		for i:=0;i<loc;i+=1{
			fmt.Printf("%d ",s[i])
		}
		
	}
	
}