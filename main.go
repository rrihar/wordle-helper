package main

import (
	"fmt"
	// "strings"
	// 	"log"
	"bufio"
	"os"
)

func check_error(e error) {
	if e != nil {
		panic(e)
	}
}

type char struct {
	data byte
	next []char
}

func find(data []char, word string) bool {
	i := 0
	for i < len(word) {

		if data == nil {
			return false
		}

		index := word[i] - 97
		fmt.Println(i, data[index].data, word[i])

		if data[index].data == word[i] {
			i++
			data = data[index].next
		} else {
			return false
		}
	}
	return true
}

func main() {
	var head [26]char

	// fmt.Println(head)
	f, err := os.Open("words_alpha.txt")

	check_error(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		word := scanner.Text()
		// fmt.Println(word)
		index := word[0] - 97
		if head[index].data == 0 {
			head[index] = char{data: word[0], next: nil}
		}

		prev_char := &head[index]

		i := 1
		for i < len(word) {
			index := word[i] - 97

			if (*prev_char).next == nil {
				(*prev_char).next = make([]char, 26)
			}

			(*prev_char).next[index].data = word[i]
			prev_char = &((*prev_char).next[index])
			i++
		}
	}

	// r = bufio.NewScanner(os.Stdin)

	// fmt.Println(head)
	toFind := "educative"

	fmt.Printf("%s found: %t", toFind, find(head[:], toFind))

}
