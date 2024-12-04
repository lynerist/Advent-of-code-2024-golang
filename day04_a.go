package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	input, _ := os.Open("input.txt")	
	sc := bufio.NewScanner(input)

	var text []string
	for sc.Scan(){
		text = append(text, sc.Text())
	}

	directions := []struct{x, y int}{{-1,0}, {-1,1}, {0,1}, {1,1}, {1,0}, {1,-1}, {0,-1}, {-1,-1}}
	var count int
	for x := range text[0]{
		for y := range text{
			if text[y][x] == 'X' {
				for _, direction := range directions{
					searchM(text, x, y, direction, &count)
				}
			}
		}
	}
	fmt.Println(count)
}

func searchM(text []string, x,y int, direction struct{x,y int}, count *int){
	if 	x+direction.x>=0 && x+direction.x<len(text[0]) && 
		y+direction.y>=0 && y+direction.y<len(text) &&
		text[y+direction.y][x+direction.x]=='M' {
		searchA(text, x+direction.x, y+direction.y, direction, count)
	}
}

func searchA(text []string, x,y int, direction struct{x,y int}, count *int){
	if 	x+direction.x>=0 && x+direction.x<len(text[0]) && 
		y+direction.y>=0 && y+direction.y<len(text) &&
		text[y+direction.y][x+direction.x]=='A' {
		searchS(text, x+direction.x, y+direction.y, direction, count)
	}
}

func searchS(text []string, x,y int, direction struct{x,y int}, count *int){
	if 	x+direction.x>=0 && x+direction.x<len(text[0]) && 
		y+direction.y>=0 && y+direction.y<len(text) &&
		text[y+direction.y][x+direction.x]=='S' {
		*count++
	}
}