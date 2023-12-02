package main

import (
	"fmt"
	"os"
	"log"
	"regexp"
	"bufio"
	"strings"
	"strconv"
)

func extractRounds(input string) []map[string]int {
	pattern := `\b(\d+)\s*(red|green|blue)\b`
	re := regexp.MustCompile(pattern)
	rounds := strings.Split(input, ";")

	var result []map[string]int

	for _, round := range rounds {
		matches := re.FindAllStringSubmatch(round, -1)
	
		// Initialize the map for the current round
		colors := make(map[string]int)
	
		for _, match := range matches {
			val, _ := strconv.Atoi(match[1])
			colors[match[2]] = val
		}
	
		// Append the map for the current round to the result slice
		result = append(result, colors)
	}
		
	return result
}


func main() {
	lims := map[string]int{"red":12, "green":13, "blue":14}
	
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list [][]map[string]int

	for scanner.Scan() {
		list = append(list, extractRounds(scanner.Text()))
	
	}

	sum := 0
	i := 0
	for _, cube_list := range list {
		i++
		good := true
		for _, cube_map := range cube_list {
			for key := range cube_map {
				if cube_map[key] > lims[key] {
					good = false
				}
			}
		}
		if good {
			sum += i
		}
	}
	fmt.Println(sum)

}