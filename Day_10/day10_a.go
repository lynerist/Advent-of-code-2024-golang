package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	input, _ := os.Open("input.txt")	
	sc := bufio.NewScanner(input)

	var topographicMap []string
	for sc.Scan() {
		topographicMap=append(topographicMap, sc.Text())
	}
	
	var score int
	for y := range topographicMap {
		for x := range topographicMap[0]{
			reached := make(map[struct{x, y int}]bool)
			if topographicMap[y][x] == '0'{
				getScore(x, y, '0'-1, topographicMap, reached) 
			}
			score += len(reached)
		}
	}
	fmt.Println(score)
}

func getScore(x,y int, from byte, topographicMap []string, reached map[struct{x,y int}]bool){
	if x<0 || y<0 || y>=len(topographicMap) || x>=len(topographicMap[0])  || 
		topographicMap[y][x] != from+1{
		return
	}
	if topographicMap[y][x] == '9'{
		reached[struct{x,y int}{x,y}] = true
		return
	}
	getScore(x-1, y, from+1, topographicMap, reached) 
	getScore(x, y+1, from+1, topographicMap, reached) 
	getScore(x+1, y, from+1, topographicMap, reached) 
	getScore(x, y-1, from+1, topographicMap, reached)
}