package main
import "fmt"

func change(number *int){

	*number = 123

}

func main(){

a :=25
p:=&a
var q *int = &a
fmt.Println("value of a: ",*p)
fmt.Println("value of p: ",p)
fmt.Println("value of q: ",q)
*p = 50


fmt.Println(a)
pp := new(int)
fmt.Println(*pp)
ss:=1
change(&ss)
fmt.Println(ss)

}
