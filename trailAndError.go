/*
 * Copyright (c) 2021
 * Author : Trinadh Koya
 * Email: trinadhkoya@gmail.com
 *
 *
 */

package main

import "fmt"

func main() {
	//arr := []int{}
	//s1 := arr[0:5] //it st
	//fmt.Println(s1[0])
	//
	//fmt.Println(cap(s1))
	var s []int

	a := append(s, 1)
	printSlice(a)

}

func printSlice(s []int) {
	fmt.Println(s)
}
