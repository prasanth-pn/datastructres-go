package main

import "fmt"

func insertion(a []int){
	for i:=1; i<len(a);i++{
		temp:=a[i]
		j:=i-1
		for j>=0&&a[j]>temp {
			a[j+1]=a[j]
			j=j-1
		}
		a[j+1]=temp

	}
	fmt.Println(a)

}

func main(){
	k:=[]int{456,23,12,22,1,23,12,45,33,43}

	insertion(k)

}