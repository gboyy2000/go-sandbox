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
		
		}
