package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

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
