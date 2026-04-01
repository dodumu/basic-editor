package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter the command you want carried out")
		scanner.Scan()
		input := scanner.Text()
		split := strings.Fields(input)
		command := split[0]
		operation := split[1]

		switch command {
		case "calc":
			if len(split) != 4 {
				fmt.Println("Not enough parameters to call the calc program")
				continue
			}
			num1, err := strconv.ParseFloat(split[2], 64)
			if err != nil {
				fmt.Println("not a valid float number")
				continue
			}
			num2, err := strconv.ParseFloat(split[3], 64)
			if err != nil {
				fmt.Println("not a valid float number")
				continue
			}
			calc(operation, num1, num2)
			continue
		case "base":
			num1 := split[2]
			convertBase(operation, num1)
		case "str":
			input := split[2:]
			w := strings.Join(input, " ")
			Format(operation, w)

		}
	}
}
