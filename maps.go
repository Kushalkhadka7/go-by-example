package main

import "fmt"

func main() {
	var newMap map[string]string = make(map[string]string)

	newMap["hello"] = "world"

	fmt.Println(newMap["hello"])

	x := map[string]int{
		"data": 10,
	}

	v, ok := x["data"]

	if !ok {
		fmt.Println("data not found")
	} else {
		fmt.Printf("data found: %v \n", v)
	}
}
