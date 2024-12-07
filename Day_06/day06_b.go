package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	x int 
	y int
}

func main(){
	input, _ := os.Open("input.txt")	
	sc := bufio.NewScanner(input)

	obstacles := make(map[Coordinate]bool)
	var guardStartPosition Coordinate

	movements := []struct{x, y int}{{0,-1},{1,0},{0,1},{-1,0}}
	currentDirection := 0

	var mapHeight, mapWidth int 
	for y:=0; sc.Scan(); y++{
		for x, point := range sc.Text(){
			if point == '^' {
				guardStartPosition = Coordinate{x,y}
			}
			if point == '#'{
				obstacles[Coordinate{x,y}] = true
			}
		}
		mapHeight++
		mapWidth = len(sc.Text())
	}

	var cycles int
	for y:=0; y<mapHeight; y++{
		for x:=0; x<mapWidth; x++{
			guard := guardStartPosition
			currentDirection = 0
			if !obstacles[Coordinate{x,y}] && (guard.x!=x || guard.y!=y){
				addedObstacle := Coordinate{x,y}
				obstacles[addedObstacle] = true
				visited := make(map[Coordinate]bool)
				var stepInLoop int
				for guard.x >= 0 && guard.x < mapWidth &&
					guard.y >= 0 && guard.y < mapHeight {
					if visited[Coordinate{guard.x, guard.y}] {
						stepInLoop++
						if stepInLoop > 100 {
							cycles++
							break
						}
					}else{
						stepInLoop=0
					}

					visited[Coordinate{guard.x, guard.y}] = true
					for obstacles[Coordinate{guard.x+movements[currentDirection].x, guard.y+movements[currentDirection].y}] {
						currentDirection = (currentDirection+1)%4
					}
					guard.x += movements[currentDirection].x
					guard.y += movements[currentDirection].y
				}
				obstacles[addedObstacle]=false
			}
		}
	}
		
	fmt.Println(cycles)
}
