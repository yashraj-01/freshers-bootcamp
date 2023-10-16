/*
Exercise link: https://docs.google.com/document/d/1LZ7q3zF9FmQNb2jGD78v5Az5qUtexcFxHWDCTqgtlRA/edit#heading=h.r2aa02zg243u
Problem 1:
*/
package main

import "fmt"

/*
Calculates the frequency of each letter (a-z) in the provided string.
Accepts a pointer to a string and a channel as input.
Stores the frequency of each letter in an int32 slice (frequency map) and sends it to the channel.
*/
func calcFreq(s *string, channel chan []int32) {
	freq := make([]int32, 26)
	for _, ch := range *s {
		freq[ch-'a']++
	}
	channel <- freq
}

func main() {
	arr := []string{"quick", "brown", "fox", "lazy", "dog"}
	l := len(arr)
	freqChannel := make(chan []int32)
	for i := 0; i < l; i++ {
		go func(s string) {
			calcFreq(&s, freqChannel)
		}(arr[i])
	}
	res := make([]int32, 26)
	for i := 0; i < l; i++ {
		temp := <-freqChannel
		for j := 0; j < 26; j++ {
			res[j] += temp[j]
		}
	}
	fmt.Println(res)
}
