package main

import "fmt"



type calculator interface{
getSpeed() int

}

type Car struct{
	name string
	speed int
}

func (c Car) getSpeed() int {
	return c.speed
}

func print(c calculator){
	fmt.Println(c.getSpeed())
}

func main(){

car1 := Car{
	name: "BMW",
	speed: 220,
}

speed := car1.getSpeed()
fmt.Println(speed)


print(car1)


}