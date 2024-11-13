package main

import (
	"fmt"
	"strings"
)

func main() {
	trim := strings.Trim("Syahrul ", " ")
	lowerCase := strings.ToLower("Syahrul Safarila")
	contains := strings.Contains("Syahrul Safarila", "Saf")
	replaceAll := strings.ReplaceAll("Syahrul Safarila", "S", "X")

	fmt.Println(trim)
	fmt.Println(lowerCase)
	fmt.Println(contains)
	fmt.Println(replaceAll)
}
