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
	evolvedStones := make(map[string]int)

	var res int
	for _, stone := range stones {
		res += evolveStone(stone, 75, evolvedStones)
	}
	fmt.Println(res)
	
}

func applyRule(stone string)[]string{
	switch{
	case stone == "0":
		return []string{"1"}
	case len(stone)%2 == 0:
		newStone, _ := strconv.Atoi(stone[len(stone)/2:])
		return []string{stone[:len(stone)/2], fmt.Sprint(newStone)}
	default:
		newStone, _ := strconv.Atoi(stone)
		return []string{fmt.Sprint(newStone*2024)}
	}
}

func evolveStone(stone string, blinks int, evolvedStones map[string]int)int{
	if res, ok := evolvedStones[fmt.Sprintf("%s_%d", stone, blinks)]; ok{
		return res
	}
	if blinks == 0{
		return 1
	}
	var res int
	for _, newStone := range applyRule(stone){
		res += evolveStone(newStone, blinks-1, evolvedStones)
	}
	evolvedStones[fmt.Sprintf("%s_%d", stone, blinks)] = res
	return res
}
