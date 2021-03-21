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
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println( cap(s))
}
