package main
import "fmt"


func sum(price, n int)(int,int){

	s := price*n
	c := s/(n-1)
	return s, c

}


func main(){
price, n := 78,87
c1,c2 := sum(price,n) 
_,c4 := sum(price,n)

fmt.Println(c1,c2)
fmt.Println(c4)
}
