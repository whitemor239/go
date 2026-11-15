package main
import "fmt"

func main(){

a:=make(map[string]int)

a["tyler"]=1
a["amin"]=2

fmt.Println(a)

a["yasi"]=3

for i,v := range a{

fmt.Println(i,v)

}


aa, exists := a["yasi"]

if exists {
	fmt.Println("exists: ",aa)
}else{
	fmt.Println("not exists")
}



}
