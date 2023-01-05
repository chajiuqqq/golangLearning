/*
Bubble sort
*/
package main
import "fmt"
func bubbleSort(a []int){
	swapped:=true
	for swapped{
		swapped=false
		for i:=0;i<len(a)-1;i++{
			if a[i]<a[i+1]{
				a[i],a[i+1] = a[i+1],a[i]
				swapped = true
			}
		}
	}

}

func main(){
	s := []int{12,23,3,34,51}
	fmt.Printf("before:%v\n",s)
	bubbleSort(s)
	fmt.Printf("after:%v",s)
}