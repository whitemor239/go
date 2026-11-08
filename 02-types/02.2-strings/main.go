package main
import "fmt"


func charr(s string){
	fmt.Println(s)
	for i := 0; i<len(s); i++{
		fmt.Printf("%c ", s[i])


}


}

func main(){

	
	
	var a string = "tyler"
	b := "goat"

	fmt.Println(a == b)
	charr("salam")
}
