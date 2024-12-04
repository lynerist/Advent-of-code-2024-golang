package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
	input, _ := os.Open("input.txt")
	sc := bufio.NewScanner(input)

	var left, right []int
	for sc.Scan(){
		line := strings.Fields(sc.Text())
		l, _ := strconv.Atoi(line[0])
		r, _ := strconv.Atoi(line[1])
		left = append(left, l)
		right = append(right, r)
	}
	sort.Ints(left)
	sort.Ints(right)

	var distance int
	for i := range left{
		if left[i]>right[i]{
			distance += left[i]-right[i]
		}else{
			distance += right[i]-left[i]
		}
	}
	fmt.Println(distance)
}

