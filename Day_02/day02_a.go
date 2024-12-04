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
	var safeReports int

	for sc.Scan(){
		var previousLevel int
		safe, crescent := true, false
		
		for i, level := range strings.Fields(sc.Text()){
			currentLevel, _ := strconv.Atoi(level)
			if i==1 {
				crescent = currentLevel>previousLevel
			}

			if i>0 && !(crescent && currentLevel-previousLevel>0 && currentLevel-previousLevel<4 ||
						!crescent && previousLevel-currentLevel>0 && previousLevel-currentLevel<4) {
				safe = false
				break
			}
			previousLevel = currentLevel
		}
		if safe {
			safeReports++
		}
	}
	fmt.Println(safeReports)
}

