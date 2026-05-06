package main

import (
	"fmt"
	"math/rand/v2"
)

func isIn(number int, array []int) bool {
	j := 0
	for j < len(array) {
		if number == array[j] {
			return true
		}
		j++
	}

	return false
}

func randomArray() []int {

	size := rand.IntN(10)

	var randArray []int

	k := 0

	for k < size {
		randArray = append(randArray, rand.IntN(101))
		k++
	}

	return randArray

}

func main() {

	inputNumbers := randomArray()

	var outputNumbers []int

	i := 0
	var number int

	fmt.Println("Willkommen bei Usersort! Bitte sortiere das folgende Array:")
	fmt.Println(inputNumbers)

	sorted := false

	for sorted == false {
		for i < len(inputNumbers) {
			fmt.Println("Nenne Zahl ", i+1, ":")
			fmt.Scan(&number)

			if isIn(number, inputNumbers) {
				outputNumbers = append(outputNumbers, number)
				i++
			} else {
				fmt.Println("Zahl nicht im ursprünglichen Array!")
			}
		}

		i = 0

		test := true
		for i < len(outputNumbers)-1 {
			if outputNumbers[i] > outputNumbers[i+1] {
				fmt.Println("Array nicht sortiert. Versuche es nochmal!")
				fmt.Println("Array: ", inputNumbers)

				test = false
				outputNumbers = nil
				i = 0

				break
			}

			i++
		}

		if test == true {
			fmt.Println("Das Array ist richtig sortiert. Gute Arbeit!")
			fmt.Println(outputNumbers)

			sorted = true
		}

	}

}
