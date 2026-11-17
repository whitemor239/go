package main
import "fmt"
import "test/model"





func main(){

user1 := model.User{
	FirstName: "tyler",
	LastName: "Mokhtari",
	Age : 28,
	Address : model.Address{

		City: "Tehran",
		Street: "first st.",

	},
} 
fmt.Println(user1)
fmt.Println("street of user1 is: ",user1.Street)


var user2 model.User

user2 = model.User{

FirstName: "amin",
LastName: "creator",
Age: 23,
Address: model.Address{

	City: "NY",
},

}

fmt.Println(user2.City)

}
