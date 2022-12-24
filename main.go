package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func getInput() (string, string) {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	check_error(err)

	input = strings.TrimSuffix(input, "\n")
	options := strings.Split(input, " ")

	var orange, green string

	green = options[0]

	if len(options) == 2 {
		orange = options[1]
	}

	return green, orange
}

func buildDictionary() []char {
	var head [26]char

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
	return head[:]

}

func filterWords(prospects []string, orange string) []string {

	if len(orange) == 0 {
		return prospects
	}

	var ans []string
	for i := 0; i < len(prospects); i++ {
		contain_flag := true
		for _, char := range orange {
			if !strings.Contains(prospects[i], string(char)) {
				contain_flag = false
				break
			}
		}
		if contain_flag == true {
			ans = append(ans, prospects[i])
		}
	}
	return ans
}

func main() {

	fmt.Println("Please wait while the dictionary is being loaded...")
	dictionary := buildDictionary()
	fmt.Println("Finished loading dictionary.")

	// fmt.Println("Press ctrl+c")

	for true {
		fmt.Println("\n******************")
		fmt.Printf("Enter green, orange(optional): ")

		green, orange := getInput()

		prospects := find(dictionary, green, 0)

		ans := filterWords(prospects, orange)

		if len(ans) > 0 {
			fmt.Println("Found these words: ", ans)
		} else {
			fmt.Println("Did not find any word that match your criteria")
		}
		fmt.Println("******************")
	}

}
