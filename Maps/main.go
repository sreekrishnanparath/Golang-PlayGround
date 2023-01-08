package main

import "fmt"

func main() {

	colors := map[string]int{
		"red": 1, "blue": 2,
	}

	fmt.Println(colors)

	emptyMap := make(map[string]int)

	emptyMap["new"] = 1
	fmt.Println(emptyMap)

	delete(emptyMap, "new")
	fmt.Println(emptyMap)

	for key, _ := range colors {
		fmt.Println(key)
	}


}