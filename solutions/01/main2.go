package main

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	"log"
	"strings"
	"strconv"
)

func reverse(inv_string string) string {
	retstring := ""
	for _, char := range inv_string {
		retstring = string(char) + retstring
	}
	return retstring
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numdict := map[string]string{
		"two": "2",
		"eight": "8",
		"nine": "9",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"one": "1",

	}
	sum := 0
	scanner := bufio.NewScanner(file)
	first_letter := ""
	last_letter := ""
	for scanner.Scan() {
		word := ""
		for _, char := range scanner.Text() {
			if first_letter != "" {
				break
			}
			if unicode.IsDigit(char) {
				first_letter = string(char)
				break
			}
			word += string(char)
			for key := range numdict {
				if strings.Contains(word, key) {
					first_letter = numdict[key]
				}
			}
		}
		word=""
		for _, char := range reverse(scanner.Text()) {
			if last_letter != "" {
				break
			}
			if unicode.IsDigit(char) {
				last_letter = string(char)
				break
			}
			word = string(char) + word
			for key := range numdict {
				if strings.Contains(word, key) {
					last_letter = numdict[key]
				}
			}
		}
		val, _ := strconv.Atoi(first_letter+last_letter)
		fmt.Printf("%s => %s%s\n",scanner.Text(), first_letter,last_letter)
		sum += val
		first_letter = ""
		last_letter = ""	
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println((sum))
	
}