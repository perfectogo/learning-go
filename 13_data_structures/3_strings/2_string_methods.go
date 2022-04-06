package main

import (
	"fmt"
	"strings"
)

func main() {
	compare()
	contains()
	replace()
	changeCase()
	split()
}

func compare() {
	var word1 = "hello, "
	var word2 = "world!"

	//compare
	// if word1>word2 return 1
	// if word1<word2 return -1
	// if word1==word2 return 0
	fmt.Println(strings.Compare(word1, word2))     // -1
	fmt.Println(strings.Compare(word1, "hello, ")) // 0
	fmt.Println(strings.Compare(word2, word1))     // 1
}

func contains() {
	text := "I'm a golang developer"
	word1 := "golang"
	word2 := "enginer"

	fmt.Println(strings.Contains(text, word1))
	fmt.Println(strings.Contains(text, word2))
}
func replace() {
	word := "hello"
	fmt.Println(word)

	word = strings.Replace(word, "l", "", 1)
	fmt.Println(word)
}

func changeCase() {
	sentences := "I'm a Golang Developer"

	s1 := strings.ToUpper(sentences)
	fmt.Println(s1)

	s2 := strings.ToLower(sentences)
	fmt.Println(s2)
}
func split() {
	text := "I'm a golang developer"

	splitedRext := strings.Split(text, " ")
	fmt.Println(splitedRext)
	splitedRext = append(splitedRext, "go")
	fmt.Println(len(splitedRext))
}
