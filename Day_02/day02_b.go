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
		var levels []int
		for _, level := range strings.Fields(sc.Text()){
			currentLevel, _ := strconv.Atoi(level)
			levels = append(levels, currentLevel)
		}

		if checkLevels(levels) {
			safeReports++
		}else{
			for i := range levels{
				copyLevels := make([]int, len(levels))
				copy(copyLevels, levels)
				if checkLevels(append(copyLevels[:i], copyLevels[i+1:]...)){
					safeReports++
					break
				}
			}
		}
	}
	fmt.Println(safeReports)
}

func checkLevels(levels []int)bool{
	var crescent bool
	var previousLevel int
	for i, currentLevel := range levels{
		if i==1 {
			crescent = currentLevel>previousLevel
		}

		if i>0 && !(crescent && currentLevel-previousLevel>0 && currentLevel-previousLevel<4 ||
					!crescent && previousLevel-currentLevel>0 && previousLevel-currentLevel<4) {
			return false
		}
		previousLevel = currentLevel
	}
	return true
}

