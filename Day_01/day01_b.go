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

	var left []int
	rightApparitions := make(map[int]int)
	for sc.Scan(){
		line := strings.Fields(sc.Text())
		l, _ := strconv.Atoi(line[0])
		r, _ := strconv.Atoi(line[1])
		left = append(left, l)
		rightApparitions[r]++
	}
	
	var similarity int
	for _, n := range left{
		similarity += n*rightApparitions[n]
	}
	fmt.Println(similarity)
}

