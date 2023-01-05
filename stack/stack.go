package stack
import (
	"fmt"
	"strconv"
)
type Stack struct{
	i int
	data [10]int
}
func (s *Stack) Push(k int){
	s.data[s.i]=k
	s.i++
}

func (s *Stack) Pop() int{
	s.i--
	return s.data[s.i]
}

//重载printf的输出
func (s Stack) String() (out string){
	
	for j:=0;j<s.i;j++{
		out = out +"["+strconv.Itoa(j)+":"+ strconv.Itoa(s.data[j]) + "] " 
	}
	return 
}
func main()  {
	s := new(Stack)
	s.Push(12)
	s.Push(24)
	fmt.Printf("stack: %v\n",s)
	fmt.Printf("data: %v\n",s.data)
}