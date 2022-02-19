// Connor Henderson

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var userInput string
	fortune := make(chan string)

	go fortunes(fortune)

	for {

		fmt.Println("Do you want a fortune?")
		fmt.Scanf(userInput)
		strings.ToLower(userInput)

		if userInput == "yes" {
			fortune <- "yes"
		}
	}
}

func fortunes(fortune chan string) {

	dat, err := os.ReadFile("Fortunes.txt")
	check(err)

	slices := strings.Split(string(dat), "%%")

	for {

		select {
		case x, ok := <-fortune:
			if ok {
				fmt.Println("Value %d was read.\n", x)
			}
		default:
			fmt.Println("No value ready, moving on.")
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
