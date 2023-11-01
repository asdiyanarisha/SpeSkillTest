package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetNarcissticNumber(numb int) bool {
	paramStr := strconv.Itoa(numb)

	slicesStr := strings.Split(paramStr, "")
	length := len(slicesStr)

	var sumNumb float64
	for _, x := range slicesStr {
		numConvert, _ := strconv.Atoi(x)
		res := math.Pow(float64(numConvert), float64(length))
		sumNumb += res
	}

	fmt.Println(sumNumb, numb)
	if sumNumb == float64(numb) {
		return true
	}

	return false
}

func main() {
	var number int

	flag.IntVar(&number, "n", 0, "help message for flag n")
	flag.Parse()

	result := GetNarcissticNumber(number)

	fmt.Println(result)
}
