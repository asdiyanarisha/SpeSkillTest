package main

import "fmt"

func FindNeedle(arr []string, key string) int {
	for i, val := range arr {
		if val == key {
			return i
		}
	}

	return 0
}

func main() {
	findValue := "red"
	paramArr := []string{"red", "blue", "yellow", "black", "grey"}
	index := FindNeedle(paramArr, findValue)
	fmt.Println(index)
}
