package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	input, _ := os.Open("input.txt")	
	sc := bufio.NewScanner(input)

	var sumResults int
	for sc.Scan(){
		testResult,_ := strconv.Atoi(strings.Split(sc.Text(), ":")[0])
		var numbers []int
		for _, input := range strings.Fields(strings.Split(sc.Text(), ":")[1]){
			number, _ := strconv.Atoi(input)
			numbers = append(numbers, number)
		}
		for _, result := range calculateResults(numbers[0], numbers[1:]){
			if testResult == result{
				sumResults+= result
				break
			}
		}
	}
	fmt.Println(sumResults)
}

func concat(a, b int)int{
	number, _ := strconv.Atoi(fmt.Sprint(a)+fmt.Sprint(b))
	return number
}

func calculateResults(result int, numbers []int)[]int{
	if len(numbers)==1{
		return []int{result*numbers[0], result+numbers[0], concat(result, numbers[0]) }
	}
	var results []int
	results = append(results, calculateResults(result*numbers[0], numbers[1:])...)
	results = append(results, calculateResults(result+numbers[0], numbers[1:])...)
	results = append(results, calculateResults(concat(result, numbers[0]), numbers[1:])...)
	return results
}
