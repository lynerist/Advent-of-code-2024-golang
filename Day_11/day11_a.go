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
	sc.Scan()
	stones := strings.Fields(sc.Text())

	for blink:=0; blink<25; blink++{
		var newStones []string
		for _, stone := range stones {
			switch{
			case stone == "0":
				newStones = append(newStones, "1")
			case len(stone)%2 == 0:
				newStone, _ := strconv.Atoi(stone[len(stone)/2:])
				newStones = append(newStones, stone[:len(stone)/2], fmt.Sprint(newStone))
			default:
				newStone, _ := strconv.Atoi(stone)
				newStones = append(newStones, fmt.Sprint(newStone*2024))
			}
		}
		stones = newStones
	}
	fmt.Println(len(stones))
}
