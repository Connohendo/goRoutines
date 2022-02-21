// Connor Henderson

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	var userInput string
	channel := make(chan string)

	go fortunes(channel)

	for {
		fmt.Println("Do you want a fortune?")
		fmt.Scanf("%s", &userInput)

		userInputLower := strings.ToLower(userInput)

		if userInputLower == "yes" {
			channel <- "yes"
		}
		if userInputLower == "no" {
			break
		}
		if userInputLower != "yes" && userInputLower != "no" {
			fmt.Println("Please enter yes or no")
			continue
		}
		<-channel
	}
}

func fortunes(channel chan string) {

	dat, err := os.ReadFile("Fortunes.txt")
	check(err)

	slices := strings.Split(string(dat), "%%")

	for {
		<-channel

		rand.Seed(time.Now().Unix())
		fmt.Println(slices[rand.Intn(len(slices))])

		channel <- "done"
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
