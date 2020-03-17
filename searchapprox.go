package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var results, guess, lowGuess, highGuess float32
	success := false
	count := 0

	// selsect a guess number to guess
	num := randomNum(0, 100)
	// display that number
	fmt.Printf("The number to guess: %f \n", num)
	// first guess
	guess = randomNum(0, 100)
	//set low and high guess
	lowGuess = 0
	highGuess = 100

	// loop until guessed correctly
	for !success {
		// compare
		results = compare(guess, num)
		// increase tries
		count++

		// adjust for results
		if results == -1 {
			// less than
			lowGuess = guess
			guess = randomNum(guess, highGuess)
		} else if results == 1 {
			// greater than
			highGuess = guess
			guess = randomNum(lowGuess, guess)
		} else {
			// number is in accapatble range
			success = true
			break
		}
	}

	// display results
	fmt.Printf("Number guessed: %f \nGuesses needed: %v \n", results, count)
}

func compare(guess float32, num float32) float32 {
	// check for approximation to the 5th fractial digit +/- .00002

	// guess is less than num
	if guess < num {

		if guess > (num - 0.00002) {
			// less than num but greater than num - 0.00002
			return num // in range
		} else if guess < (num - 0.00002) {
			// less than num but also less than num - 0.00002
			return -1 // to low
		}
	}

	// guess is greater than num
	if guess > num {
		if guess < (num + 0.00002) {
			// greater than num but less than num + 0.00002
			return num // in range
		} else if guess > (num + 0.00002) {
			// greater than num but also greater than num + 0.00002
			return 1 // to low
		}
	}

	// the rare occasion they are equal
	return num // equal
}

func randomNum(min float32, max float32) float32 {
	// set seed
	rand.Seed(time.Now().UnixNano())
	// generate a guess number from 0 - num
	n := min + rand.Float32()*(max-min)
	// return guess num
	return n
}
