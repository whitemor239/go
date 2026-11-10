package main
import "fmt"

func main(){

num:=4

switch num{

	case 1,2,3:
		fmt.Println(num)
//	case 2:
//		fmt.Println(num)
//	case 3:
//		fmt.Println(num)
	default:
		fmt.Println("error")

}

}
