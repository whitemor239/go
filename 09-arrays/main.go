package main
import "fmt"

func main(){

//var a [3]int

a:=[...]int{11,22,33,44}
fmt.Println("a is a array:",a)
for i:=0;i<len(a);i++{
	fmt.Println(a[i])
}

for i, v :=range a{
fmt.Println(i,v)

}

b := [2][2]string{
	
	{"a","b"},
	{"c","d"},


}
fmt.Println(b)
}
