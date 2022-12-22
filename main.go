package main

import (
	"fmt"
	// "strings"
	//  "log"
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
	end  bool
}

func find(data []char, word string, i int) []string {

	end := false
	for i < len(word) {

		if data == nil {
			return []string{}
		}
		index := word[i] - 97
		if index == 29 {
			ans := []string{}
			id := 0
			for id < 26 {
				if i == len(word)-1 {
					tmp := find(data, word[:i]+string(id+97), i)
					if len(tmp) != 0 {
						ans = append(ans, tmp...)
					}
				} else {
					tmp := find(data, word[:i]+string(id+97)+word[i+1:], i)
					if len(tmp) != 0 {
						ans = append(ans, tmp...)
					}
				}
				id++
			}
			return ans
		} else {
			// fmt.Println(i, word[i], data[index].data)
			if data[index].data == word[i] {
				i++
				end = data[index].end
				data = data[index].next
			} else {
				return []string{}
			}
		}

	}
	if end {
		return []string{word}
	} else {
		return []string{}
	}

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
		(*prev_char).end = true
	}

	// r = bufio.NewScanner(os.Stdin)

	// fmt.Println(head[0])
	toFind := "r~l~y"

	ans := find(head[:], toFind, 0)

	if len(ans) == 0 {
		fmt.Println(toFind, ": Did not find the word in the dictionary")
	} else {
		fmt.Println(toFind, "Found these words in the dictionary", ans)
	}

}
