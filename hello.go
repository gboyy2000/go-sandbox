package main

	import (
		"fmt"
		)

	func main() {

  		var name string
  		fmt.Print("Please enter your name: ")
  		fmt.Scan(&name)
  		fmt.Printf("Hello %v!\n" , name)

  		var age int
  		fmt.Print("Please enter your age: ")
  		fmt.Scan(&age)
  		fmt.Printf("Your name is %v and you are %d years old.\n" , name , age)

			newMessage := fmt.Sprintf("You have stated that you are %d years old and your name is %v." , age , name)
  		fmt.Println(newMessage)

		}
