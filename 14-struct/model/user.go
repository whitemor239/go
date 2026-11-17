package model


type User struct{
	FirstName string
	LastName string
	Age int 
	Address
}

type Address struct{
	City string
	Street string
}


