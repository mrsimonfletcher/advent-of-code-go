package main

import (
	"encoding/json"
	"io/ioutil"
)

func findNumbers(input interface{}) []int {
	numbers := []int{}

	switch input := input.(type) {
	case []interface{}:
		for _, value := range input {
			numbers = append(numbers, findNumbers(value)...)
		}
	case map[string]interface{}:
		noRed := true

		for _, value := range input {
			if str, ok := value.(string); ok && str == "red" {
				noRed = false
				break
			}
		}

		if noRed {
			for _, value := range input {
				numbers = append(numbers, findNumbers(value)...)
			}
		}
	case float64:
		numbers = append(numbers, int(input))
	}

	return numbers
}

func main() {
	input, err := ioutil.ReadFile("./day-12-input.txt")
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{}, 0)
	json.Unmarshal(input, &data)

	sum := 0
	for _, num := range findNumbers(data) {
		sum += num
	}

	println(sum)
}
