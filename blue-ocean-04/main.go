package main

import "fmt"

func BlueOcean(arr []int, key []int) []int {
	var newArr []int
	mapFind := make(map[int]bool)

	for _, val := range key {
		ok, _ := mapFind[val]
		if !ok {
			mapFind[val] = true
		}
	}

	for _, numb := range arr {
		ok, _ := mapFind[numb]
		if !ok {
			newArr = append(newArr, numb)
		}
	}

	return newArr
}

func main() {
	paramDiffer := []int{5}
	//paramArr := []int{1, 2, 3, 4, 6, 10}
	paramArr := []int{1, 5, 5, 5, 5, 3}
	index := BlueOcean(paramArr, paramDiffer)
	fmt.Println(index)
}
