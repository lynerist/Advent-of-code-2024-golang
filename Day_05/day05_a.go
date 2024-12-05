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

	pagesThatMustFollow := make(map[string][]string)

	for sc.Scan() && sc.Text()!="" {
		pages := strings.Split(sc.Text(),"|")
		pagesThatMustFollow[pages[0]] = append(pagesThatMustFollow[pages[0]], pages[1])
	}

	var sumOfMiddlePages int
	for sc.Scan() {
		correctOrder := true
		alreadyPrinted := make(map[string]bool)
		pages := strings.Split(sc.Text(), ",")
		for _, page := range pages{
			for _, pageThatMustFollow := range pagesThatMustFollow[page]{
				if alreadyPrinted[pageThatMustFollow]{
					correctOrder = false
				}
			}
			alreadyPrinted[page] = true
		}
		if correctOrder {
			middlePage, _ := strconv.Atoi(pages[len(pages)/2])
			sumOfMiddlePages += middlePage
		}
	}
	fmt.Println(sumOfMiddlePages)
}
