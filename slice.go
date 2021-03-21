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
	"sort"
	"strconv"
	"strings"
)

func main() {
	mSlice := make([]int, 0,3)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter an integer (X to exit):")

	for input.Scan() {
		givenInput := input.Text()
		if strings.ToLower(givenInput) == "x" {
			os.Exit(1)
		}
		val, err := strconv.Atoi(givenInput)
		if err != nil {
			fmt.Println("Looks like we have encountered something other than string ")
		} else {
			mSlice = append(mSlice, val)
			sort.Ints(mSlice[:])
			fmt.Println(mSlice)
		}
		fmt.Println("Please enter an integer (X to exit):")
	}

}
