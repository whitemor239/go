package main

import (
	"fmt"
	"time"
)

func hello(data chan int){
	time.Sleep(4000 * time.Millisecond)
	fmt.Println("waited 4 sec: ",1)
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("waited 3 sec: ",2)
	data <- 333	
}

func cal1(number int, cal chan int){
	number *= 2
	cal <-number
}

func cal2(number int, cal chan int){
	number += 2
	cal <-number
}


func getnumber(data chan int){
	for i := 0; i < 10; i++ {
		data<- i
	}
	close(data)
}
// func sender(inti int)int{
// 	return inti
// }





func main(){
	// go hello()
	// var a chan int
	// fmt.Println(a)	
	// b := make(chan int)
	// fmt.Println(b)
	// data := 123
	// b <- data
	// fmt.Printf("%T",b)

	messages := make(chan string)
	go func(){
		messages <- "hellow from a goroutines"
	}()
	msg := <-messages
	fmt.Println(msg)

	ch := make(chan int)
	go func(){
		ch <-42
		fmt.Println("value sent")
	}()
	inti := <-ch
	fmt.Println("value recieved:",inti)


	data := make(chan int)

	go hello(data)
	<-data
	
	number:=10
	a:= make(chan int)
	b:= make(chan int)

	go cal1(number, a)
	go cal2(number, b)
	fmt.Println("main")
	
	s1, s2:= <-a, <-b
	fmt.Println(s1+s2)
	num:=make(chan int)
	go getnumber(num)
	// for{
	// 	v, ok:= <-num
	// 	if ok== false{
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }
	for v:= range num{
		fmt.Println(v)
	}
}