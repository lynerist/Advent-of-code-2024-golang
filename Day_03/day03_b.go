package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main(){
	input, _ := os.Open("input.txt")
	
	var total int
	rexp := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\)`)
	sc := bufio.NewScanner(input)
	enabled := true
	for sc.Scan(){
		matches := rexp.FindAllStringSubmatch(sc.Text(), -1)
		for _, match := range matches {
			switch match[0]{
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if enabled {
					X,_ := strconv.Atoi(match[1])
					Y,_ := strconv.Atoi(match[2])
					total += X*Y
				}
			}
			fmt.Println(match, enabled)
		}
	}
	fmt.Println(total)
}
