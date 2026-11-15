package main
import "fmt"
//import "make"

func change(numbers []int){

for i:=range numbers{
	
	numbers[i]-=1

}

}

func main(){

a :=[6]int{0,1,2,3,4,5}
var b []int=a[2:4]
fmt.Println("before: ",b)

for i := range b {

	b[i]++

}
fmt.Println("after: ",b)

i := make([]int, 3, 5)
i = append(i,100,200,300,400)
fmt.Println("i before is: ",i)

change(i)
fmt.Println("i after is: ",i)
}
