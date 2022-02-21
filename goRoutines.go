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
	// channel creation
	channel := make(chan string)

	//create go routine for fortunes function
	go fortunes(channel)

	//enter forever loop
	for {
		//prompt user and get input from terminal
		fmt.Println("Do you want a fortune?")
		fmt.Scanf("%s", &userInput)

		// cast to lower case to accept all capitalization
		userInputLower := strings.ToLower(userInput)

		// if yes send a message through the channel, if no end the main routine,
		//if anything other than yes or no jump back to the top of the for loop
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
		// wait for a message from the channel (a fortune)
		<-channel
	}
}

func fortunes(channel chan string) {

	// read a file into a string and check for errors
	dat, err := os.ReadFile("Fortunes.txt")
	check(err)

	// create a slice with the read file
	slices := strings.Split(string(dat), "%%")

	// enter a forever loop
	for {
		// wait for a message from the channel
		<-channel

		// random pick an index of slice and print it
		rand.Seed(time.Now().Unix())
		fmt.Println(slices[rand.Intn(len(slices))])

		// send a message down the channel to main
		channel <- "done"
	}
}

// error check function
func check(e error) {
	if e != nil {
		panic(e)
	}
}
