package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 	var names = []string{"John", "Wick"}
	// 	printMessage("hallo", names)
	// }

	// func printMessage(message string, arr []string) {
	// 	var nameString = strings.Join(arr, " ")
	// 	fmt.Println(message, nameString)

	rand.Seed(time.Now().Unix())

	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number : ", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number : ", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number : ", randomValue)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}
