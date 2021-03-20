/*
 * Copyright (c) 2021
 * Author : Trinadh Koya
 * Email: trinadhkoya@gmail.com
 *
 *
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("enter a string to find your IAN")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		givenInput := input.Text()
		if strings.ToLower(givenInput) == "q" {
			os.Exit(1)
		}
		if !foundIAN(strings.ToLower(givenInput)) {
			fmt.Println("Not Found!")
		} else {
			fmt.Println("Found!")
		}
		fmt.Println("Enter another value to truncate or type Q to exit")

	}

}
func foundIAN(s string) bool {
	trimmedString := strings.ReplaceAll(s, " ", "")
	if strings.HasPrefix(trimmedString, "i") && strings.HasSuffix(trimmedString, "n") && strings.Contains(trimmedString, "a") {
		return true
	}
	return false
}
