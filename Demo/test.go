package main

import "fmt"


type Contact struct {
	greeting string
	name     string
}

func Greet(contact Contact) {
	myGreetingMas, myNameMas := CreateMessage(contact.name, contact.greeting, "howdy","Marcus")
	fmt.Print(myGreetingMas)
	fmt.Print(myNameMas)
}

// VARIADIC FUNCTIONS
// a variable number of parameters of a certain type
// useful when we don't know how many parameters are going to be passed in
func CreateMessage(name string, greeting ...string) (myGreeting string, myName string) {
	// change the index from 0 to 1 and watch what happens
	myGreeting = greeting[2] + " " + name
	myName = "\n Hey " + name
	return myGreeting, myName
}

func main() {

	var t = Contact{"Good to see you,", "Tim"}
	Greet(t)

	u := Contact{"Glad you're in class,", "Jenny"}
	Greet(u)

	v := Contact{}
	v.greeting = "We're learning great things,"
	v.name = "Julian"
	Greet(v)
}