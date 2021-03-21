/*
 * Copyright (c) 2021
 * Author : Trinadh Koya
 * Email: trinadhkoya@gmail.com
 *
 *
 */

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	m := make(map[string]string)
	fmt.Println("enter your name")
	fmt.Scanln(&name)
	fmt.Println("enter your address")
	fmt.Scanln(&address)
	m["name"] = name
	m["address"] = address
	res, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error",err)
	}
	fmt.Println(string(res))

}
