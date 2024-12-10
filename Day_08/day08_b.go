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

	antennas := make(map[rune][]Coordinate)
	antinodes := make(map[Coordinate]bool)

	var mapHeight, mapWidth int 
	for y:=0; sc.Scan(); y++{
		for x, signal := range sc.Text(){
			if signal != '.' {
				for _, position := range antennas[signal] {
					for harmonic := 0; harmonic <100; harmonic++{
						antinodes[Coordinate{x + (x-position.x)*harmonic, y + (y-position.y)*harmonic}] = true
						antinodes[Coordinate{position.x + (position.x-x)*harmonic, position.y + (position.y-y)*harmonic}] = true
					}
				}
				antennas[signal] = append(antennas[signal], Coordinate{x,y})
			}
		}
		mapHeight++
		mapWidth = len(sc.Text())
	}

	toDelete := make(map[Coordinate]bool)
	for antinode := range antinodes {
		if antinode.x<0 || antinode.y<0 || antinode.x >= mapWidth || antinode.y >= mapHeight {
			toDelete[antinode]=true
		}
	}
	
	for antinode := range toDelete {
		delete(antinodes, antinode)
	}

	fmt.Println(len(antinodes))
}