package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Exercise link: https://docs.google.com/document/d/1LZ7q3zF9FmQNb2jGD78v5Az5qUtexcFxHWDCTqgtlRA/edit#heading=h.r2aa02zg243u
// Problem 2:

// Simulates a feedback mechanism with a random delay.
// Generates and returns a random integer between 0 and 9 after a random time delay.
func getFeedback() int {
	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	res := rand.Intn(10)
	fmt.Println(res)
	return res
}

func main() {
	noOfStudents := 200
	totalRatings := 0

	var wg sync.WaitGroup

	for i := 0; i < noOfStudents; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			totalRatings += getFeedback()
		}()
	}

	wg.Wait()

	avgRating := totalRatings / noOfStudents
	fmt.Println("Avg Rating:", avgRating)
}
