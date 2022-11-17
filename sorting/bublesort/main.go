package main

import (
	"fmt"
	//"reflect"
)



func buble( a []int){
	
	//b:=reflect.Swapper(a)

for i := 0; i <len(a)-1; i++ {
for	j:=0; j<(len(a)-i)-1; j++ {
		if a[j]>a[j+1]{
		temp:=a[j]
		a[j]=a[j+1]
		a[j+1]=temp
		}

	}
	
}
}

func main(){
	k:=[]int{10,23,1,34,22,12,567,234,13}
	buble(k)
	fmt.Println(k)

}