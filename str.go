package main

import (
	"fmt"
	"strings"
)

func Format(operation string, every string) {
	switch operation {
	case "upper":
		fmt.Println(strings.ToUpper(every))
	case "lower":
		fmt.Println(strings.ToLower(every))
	case "cap":
		split := strings.Fields(every)
		for i := 0; i < len(split); i++ {
			split[i] = strings.Title(split[i])
		}
		fmt.Println(strings.Join(split, " "))
	case "title":
		split := strings.Fields(every)
		for i := 0; i < len(split); i++ {
			if len(split[i]) >= 3 {
				split[i] = strings.Title(split[i])

			} else if len(split[0]) < 3 {
				split[0] = strings.Title(split[0])
			}

		}
		fmt.Println(strings.Join(split, " "))
	case "snake":
		every = strings.ReplaceAll(every, " ", "_")
		fmt.Println(every)
	case "reverse":
		split := strings.Fields(every)
		for i := 0; i < len(split); i++ {
			word := split[i]
			r := []rune(word)
			for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
				r[i], r[j] = r[j], r[i]
			}
			var abeg string
			abeg += string(r)

			fmt.Printf(abeg + " ")
		}
		fmt.Println()
	}

}
