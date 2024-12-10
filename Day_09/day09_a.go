package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	input, _ := os.Open("input.txt")	
	sc := bufio.NewScanner(input)

	sc.Scan()
	var memory *node
	var lastNode *node
	for i, r := range sc.Text(){
		if i==0{
			memory = &node{false, i/2, toInt(r), nil, nil}
			lastNode = memory
		}else if i%2==0 {
			lastNode.next = &node{false, i/2, toInt(r), nil, lastNode}
			lastNode = lastNode.next
		}else{
			lastNode.next = &node{true, -1, toInt(r), nil, lastNode}
			lastNode = lastNode.next
		}
	}
	
	tmp := memory
	
	for tmp != nil{
		if !tmp.isSpace{
			tmp = tmp.next
			continue
		}
		for lastNode.isSpace || lastNode.len == 0{
			lastNode = lastNode.prev
			lastNode.next = nil
		}
		tmp.isSpace = false
		tmp.id = lastNode.id
		if tmp.len > lastNode.len {
			tmp.next = &node{true, -1, tmp.len-lastNode.len, tmp.next, tmp}
			tmp.len = lastNode.len

			lastNode = lastNode.prev
			lastNode.next = nil
		}else{
			lastNode.len -= tmp.len
		}
		tmp = tmp.next
	}
	var checksum,position int
	for tmp = memory; tmp != nil; tmp = tmp.next {
		for i:=0; i<tmp.len; i++{
			checksum += tmp.id*position
			position++
		}
	}
	fmt.Println(checksum)
}

type node struct{
	isSpace bool
	id int
	len int
	next *node
	prev *node
}

func toInt(r rune)int{
	return int(r)-48
}
