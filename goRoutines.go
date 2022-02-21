// Connor Henderson

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var userInput string
	fortune := make(chan string)

	go fortunes(fortune)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Do you want a fortune?")
		userInput, _ := reader.ReadString('\n')
		strings.ToLower(userInput)
		fmt.Println("here is your user input: " + userInput)

		if userInput == "yes" {
			fmt.Println("you made it to yes")
			//fortune <- "yes"
		}
		if userInput == "no" {
			return
		}
	}
}

func fortunes(fortune chan string) {

	dat, err := os.ReadFile("Fortunes.txt")
	check(err)

	slices := strings.Split(string(dat), "%%")

	for {
		<-fortune
		fmt.Println(slices[1])
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
