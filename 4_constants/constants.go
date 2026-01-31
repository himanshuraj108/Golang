package main

import "fmt"

// const age = 30

// := we cant use this outside the main fun but can use other way

func main() {
	// const name string = "GoLang"
	// fmt.Println(name)
	// fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
