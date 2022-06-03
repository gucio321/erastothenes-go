package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	var maxNumber string
	fmt.Print("Enter number of number you want to calculate: ")
	fmt.Scanln(&maxNumber)
	maxNumberInt, err := strconv.Atoi(maxNumber)
	if err != nil {
		log.Fatal("error converting your answer to intager. Are you sure you know what is this algorithm used for?")
	}

	t := time.Now()
	result := getPrimeNumbers(maxNumberInt)
	duration := time.Now().Sub(t)
	fmt.Println(result)
	fmt.Printf("duration was %v\n", duration)
}

func getPrimeNumbers(max int) []int32 {
	list := make([]int32, 0)
	for i := int32(2); i <= int32(max); i++ {
		list = append(list, i)
	}

	currentIndex := 0
	lastIndex := list[len(list)-1]
	for {
		multiple := list[currentIndex]
		if multiple*2 > lastIndex {
			break
		}

		for i := currentIndex + 1; i < len(list); i++ {
			if list[i]%multiple == 0 {
				list = append(list[:i], list[i+1:]...)
			}
		}

		currentIndex++
	}

	return list
}
