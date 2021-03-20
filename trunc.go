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
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Enter value to truncate")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		givenInput := input.Text()
		if strings.ToLower(givenInput) == "q" {
			os.Exit(1)
		}
		if IsLetter(givenInput) {
			fmt.Println("Looks like we have encountered different format ")
			break
		}
		truncatedValue, _ := strconv.ParseFloat(givenInput, 8)
		res := int(truncatedValue) //Convert uint64 To int
		fmt.Printf("Truncated value is: %d\n", res)
		fmt.Println("Enter another value to truncate or type Q to exit")
	}

}
func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
