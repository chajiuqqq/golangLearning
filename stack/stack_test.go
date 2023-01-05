package stack
import "testing"
func TestPushAndPop(t *testing.T){
	s := new(Stack)
	s.Push(12)
	if s.Pop()!=12{
		t.Log("pop没给出12")
		t.Fail()
		
	}
	
}
