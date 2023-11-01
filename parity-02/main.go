package main

import "fmt"

func ParityOutlier(arr []int) int {
	var oddNumber []int
	var evenNumber []int

	for _, val := range arr {
		if val%2 == 0 {
			evenNumber = append(evenNumber, val)
		} else {
			oddNumber = append(oddNumber, val)
		}
	}

	if len(oddNumber) == 1 {
		return oddNumber[0]
	}

	if len(evenNumber) == 1 {
		return evenNumber[0]
	}

	return 0
}

func main() {
	// dummy parameters
	//arrData := []int{2, 4, 0, 100, 4, 11, 2602, 36}
	//arrData := []int{160, 3, 1719, 19, 11, 13, -21}
	arrData := []int{11, 13, 15, 19, 9, 13, -21}

	res := ParityOutlier(arrData)
	if res > 0 {
		fmt.Println(res)
	} else {
		fmt.Println("false")
	}
}
