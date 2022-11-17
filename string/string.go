package main

import (
	"fmt"
	"strings"
)


func main(){
	str:=`pra`
	str1:=`prathap pra santh`
	str2:=`prasanth`
	fmt.Printf("%c\n",str[2])//single characterprinting
	name:= str +""+ str
	fmt.Println(name)
	fmt.Println(strings.Compare(str,str1))
	fmt.Println(strings.Compare(str,str2))
	result:=strings.Contains(str1,str)
	fmt.Println(result)
	fmt.Println(strings.Replace(str,"pra","kyle",1))
	fmt.Println(strings.Replace("frogram","f","p",1))
	fmt.Println(strings.ToLower("HI HOW are you"))
	fmt.Println(strings.ToUpper("prasanthpn"))
	fmt.Println(strings.Split())
	


}