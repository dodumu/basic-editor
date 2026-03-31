package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertBase(operation string, num1 string) {
	switch operation {
	case "hex":
		value, err := strconv.ParseInt(num1, 16, 64)
		if err == nil {
			fmt.Println("Your output is:", value)
		} else {
			fmt.Println("Not a hex value")
		}
	case "bin":
		value, err := strconv.ParseInt(num1, 2, 64)
		if err == nil {
			fmt.Println("Your output is:", value)
		} else {
			fmt.Println("Not a bin value")
		}
	case "dec":
		a, _ := strconv.Atoi(num1)
		binval := strconv.FormatInt(int64(a), 2)
		hexval := strconv.FormatInt(int64(a), 16)
		fmt.Printf("Your results are:\n Bin: %v\n Hex: %v\n", binval, strings.ToUpper(hexval))
	default:
		fmt.Printf("Missing base\n Very vital info")

	}
}
