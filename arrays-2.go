package main

import "fmt"

func main(){
	b := [3]string{"anak", "bos", "kapal"}
	fmt.Println(b)

	c := [...]string{"melda", "nophia"}
	fmt.Println(c)
	fmt.Println(len(c))
}