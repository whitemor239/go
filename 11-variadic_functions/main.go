package main

import "fmt"


func sum(nums ...int){

	fmt.Print(nums, " " )
	total := 0

for _,num :=range nums {
total += num
}
fmt.Println(total)
}

func find(number int, numb ...int){

found:=false

for i,v:= range numb{

	if v==number{

		fmt.Println("the number has found in index: ",i)
		found =true
	}
	
}
if !found{
	fmt.Println("not found")
}





}

func main(){


sum(1,2)
sum (1,2,3,4,5,6,7,8,9)
find(10,11,12,14,10)


}
