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

	directions := []struct{x, y int}{{-1,1},{1,1},{1,-1},{-1,-1}}
	var count int
	for x := range text[0]{
		for y := range text{
			if text[y][x] == 'M' {
				for _, direction := range directions{
					searchA(text, x+direction.x, y+direction.y, direction, &count)
				}
			}
		}
	}
	fmt.Println(count/2)
}

func searchA(text []string, x,y int, direction struct{x,y int}, count *int){
	if 	x>=0 && x<len(text[0]) && y>=0 && y<len(text) &&
		text[y][x]=='A' {
		if searchS(text, x+direction.x, y+direction.y){
			if searchM(text, x-direction.x, y+direction.y) && searchS(text, x+direction.x, y-direction.y) ||
				searchM(text, x+direction.x, y-direction.y) && searchS(text, x-direction.x, y+direction.y){
					*count++
			}
		}
	}
}

func searchS(text []string, x,y int)bool{
	return x>=0 && x<len(text[0]) && y>=0 && y<len(text) && text[y][x]=='S'
}

func searchM(text []string, x,y int)bool{
	return x>=0 && x<len(text[0]) && y>=0 && y<len(text) && text[y][x]=='M'
}
