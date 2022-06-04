package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	auth()
}

func auth() {
	var username string
	var password string
	fmt.Println("Username: ")
	fmt.Scan(&username)
	fmt.Println("Password: ")
	fmt.Scan(&password)

	fmt.Printf("Welcome %v.", username)
}
