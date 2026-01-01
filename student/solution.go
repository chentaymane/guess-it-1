package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := []int{}
	first := true
	
	for scanner.Scan() {
		line := scanner.Text()
		input, _ := strconv.Atoi(line)
		
		if first {
			numbers = append(numbers, input)
			first = false
			continue
		}
		
		
		if len(numbers) > 0 {
			last := numbers[len(numbers)-1]
			prediction := (last + input) / 2
			fmt.Println(prediction-25, prediction+25)
		}
		
		numbers = append(numbers, input)
	}
}