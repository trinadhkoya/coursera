package main

import "fmt"

func main()  {

	var floatInput float32

	fmt.Println("Please input a float:")
	fmt.Scan(&floatInput)

	fmt.Printf("The truncated value is: %.0f\n", floatInput)
}
