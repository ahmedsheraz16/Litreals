package main

import (
	"fmt"
)

func main() {
	newName := []string{}
	names := []string{"Ahmed Saleem", "Ahmed Sheraz", "Hamza Anis", "Zain Bhatti"}

	for _, s := range names {

		n := 0

		for i, z := range s {
			if z == ' ' {
				n = i
				break
			}
		}
		shortName := string(s[0]) + string(s[n+1])
		newName = append(newName, shortName)
	}
	fmt.Print("The Names are", newName)
}
