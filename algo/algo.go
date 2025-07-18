//package main

// import (
// 	"fmt"
// 	"sort"
// 	"strings"
// )

//import "fmt"

// func algo() {

// 	questionOne()
// 	questionTwo()
// 	questionThree()

// }

// /*  Write a function that takes a slice of integers and returns the count of even numbers in it. */

// func questionOne() {
// 	counter := 0
// 	fmt.Printf("Amount of even numbers: %d\n", evenChecker(counter))
// }

// func evenChecker(counter int) int {
// 	numbers := []int{1, 2, 3, 4, 5, 6}

// 	for _, number := range numbers {

// 		if number%2 == 0 {
// 			counter++
// 		}
// 	}
// 	return counter
// }

// /* Given a slice of strings, return a map showing how many times each word appears. */

// func questionTwo() {
// 	slice := []string{"go", "is", "fun", "go", "go", "fun"}

// 	wordChecker := make(map[string]int)

// 	for _, word := range slice {
// 		wordChecker[word]++
// 	}

// 	fmt.Println(wordChecker)
// }

// /* 3. Write a function that checks if two strings are anagrams of each other. */

// func questionThree() {

// 	fmt.Println(wordChecker("listen", "silent"))

// }

// func wordConverter(s string) string {
// 	//split the words into letters and sort them alphabetically then put them back into one word
// 	letters := strings.Split(s, "") //splits letters in the word into an array

// 	sort.Strings(letters) //orders the letters alphabetically in the new array

// 	return strings.Join(letters, "") //joins the letters back into one word

// }

// func wordChecker(wordOne, wordTwo string) bool {
// 	//converting to array by making the variables rune type
// 	return wordConverter(wordOne) == wordConverter(wordTwo)
// }