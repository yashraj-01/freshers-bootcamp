package main

import "fmt"

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
