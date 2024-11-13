package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("1000")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resultInt)
	}

	stringInt := strconv.Itoa(10)

	fmt.Println(stringInt)
}
