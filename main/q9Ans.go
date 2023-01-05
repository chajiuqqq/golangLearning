package main
import (
	"fmt"
	"strconv"
)
type stack struct{
	i int
	data [10]int
}
func (s *stack) push(k int){
	s.data[s.i]=k
	s.i++
}

func (s *stack) pop() int{
	s.i--
	return s.data[s.i]
}

//重载printf的输出
func (s stack) String() (out string){
	
	for j:=0;j<s.i;j++{
		out = out +"["+strconv.Itoa(j)+":"+ strconv.Itoa(s.data[j]) + "] " 
	}
	return 
}
func main()  {
	s := new(stack)
	s.push(12)
	s.push(24)
	fmt.Printf("stack: %v\n",s)
	fmt.Printf("data: %v\n",s.data)
}