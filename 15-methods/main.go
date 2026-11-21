package main

import "fmt"


type User struct{
	Age int
	Name string
	Address
}

func (u User) print(){
	fmt.Println(u.Name,u.Age)
}

func (u *User) changeName(newName string){
	u.Name = newName
}

func (u *User) changeAge(newAge int){
	u.Age = newAge
}

type Address struct{
	code int
}

func (a Address) printCode(){
	fmt.Println("code: ",a.code)
}

func main() {

	user := User{

		Age : 22,
		Name: "Tyler",
		Address: Address {
			123,
		},

	}

	user.print()
	user.changeName("amin")
	user.changeAge(23)
	user.print()
	user.printCode()

}
