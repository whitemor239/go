package main
import "fmt"

func main(){

fmt.Println("LOOPS IN GO")

for i:=1;i<5;i++{
if i==3{
continue
}
fmt.Println(i)
}
}
