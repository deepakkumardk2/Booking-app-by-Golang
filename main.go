package main
import "fmt"	
func main() {
	fmt.Println("Welcome to our conference booking application")
	fmt.Println("Get your ticket here to attend the conference")

	conferenceName1 := "Go Conference"
	fmt.Println("Conference Name: ", conferenceName1)

	var conferenceName2 string = "Go Conference"
	fmt.Println("Conference Name: ", conferenceName2)

	const conferenceName3 = 50
	fmt.Println("Conference Name: ", conferenceName3)

	name := "Go Conference"
	fmt.Println("Conference Name: ", name) 
	//golang is a very powerful language
	// i want some more details about the conference
	// i want to became master in git and github
	var username string ;
	fmt.Println("Enter your name: ")
	fmt.Scanln(&username)
	fmt.Println("Username: ", username)
	
	
}

