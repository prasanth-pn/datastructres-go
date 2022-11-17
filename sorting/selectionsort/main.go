package main

import "fmt"

func selectionsort(a []int){
	for i := 0; i < len(a); i++ {
		min:=i
		for j:=i+1;j<len(a);j++{
			if a[j]<a[min]{
			min=j
		}
	}
		if min!=i{
			temp:=a[i]
			a[i]=a[min]
			a[min]=temp
		}
	}
}
func main(){
	k:=[]int {345,2,12,23,56,34,231,1,2,23}
	selectionsort(k)
	fmt.Println(k)
	fmt.Println(k)

}