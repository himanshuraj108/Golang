package main

import "fmt"

// for -> only construct in go for looping
func main()  {
	// while loop - there is no while loop in Golang
	// i := 1
	// for i <= 3{
	// 	fmt.Println(i)
	// 	i++ // i = i + 1
	// }

	// infinite loop
	// for{
	// 	println("1")
	// }

	// classic for loop
	// for i := 0; i<3; i++{
	// 	 fmt.Println(i)
	// }

	// for i:=1;i<=10;i++{
	// 	if i == 2{
	// 		continue 
	// 	}
	// 	if i == 8{
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	for i:= range 3{
		fmt.Println(i)
	}
	
}