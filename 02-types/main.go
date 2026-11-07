package main

import "fmt"
import "unsafe"

func main(){

	var name string = "tyler"
	var x bool = true
	var y bool = true
	z := x && y
	fmt.Println(name)
	fmt.Println(z) 
	
	var a = 10
	b := 90
	
	fmt.Println("the value of a is ",a,"and the b is",b)
	fmt.Printf("type of a is %T, size of a is %d bytes\n",a,unsafe.Sizeof(a))
	
	var flt float64 = 1.12

	fmt.Println(flt)

	c := 8 + 2i
	d := complex(8,2)
	fmt.Println(c,d)

}
