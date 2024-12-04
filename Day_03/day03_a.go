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
	rexp := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	sc := bufio.NewScanner(input)
	for sc.Scan(){
		matches := rexp.FindAllStringSubmatch(sc.Text(), -1)
		for _, match := range matches {
			X,_ := strconv.Atoi(match[1])
			Y,_ := strconv.Atoi(match[2])
			total += X*Y
		}
	}
	fmt.Println(total)
}

