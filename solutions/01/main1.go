package main

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	"log"
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
	file, err := os.Open("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	first_letter := ""
	last_letter := ""
	for scanner.Scan() {
		index_of_first := 0
		for _, char := range scanner.Text() {
			index_of_first++
			if unicode.IsDigit(char) {
				first_letter = string(char)
				break
			}
		}
		index_of_last := len(scanner.Text())
		for _, char := range reverse(scanner.Text()) {
			index_of_last--
			if unicode.IsDigit(char) {
				last_letter = string(char)
				break
			}
			if index_of_first > index_of_last {
				break
			}
		}
		val, _ := strconv.Atoi(first_letter+last_letter)
		fmt.Println(first_letter+last_letter)
		sum += val
		first_letter = ""
		last_letter = ""	
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println((sum))
	
}