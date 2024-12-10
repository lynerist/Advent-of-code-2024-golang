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
			lastNode.next = &node{true, 0, toInt(r), nil, lastNode}
			lastNode = lastNode.next
		}
	}
	
	tmp := lastNode
	
	for tmp != nil{
		if tmp.isSpace{
			tmp = tmp.prev
			continue
		}

		for possibleSpace := memory; possibleSpace != tmp; possibleSpace=possibleSpace.next{
			if possibleSpace.isSpace && possibleSpace.len >= tmp.len {
				possibleSpace.id = tmp.id
				possibleSpace.isSpace = false
				if possibleSpace.len>tmp.len{
					possibleSpace.next = &node{true, 0, possibleSpace.len-tmp.len, possibleSpace.next, possibleSpace}
					possibleSpace.next.prev = possibleSpace.next

					possibleSpace.len = tmp.len
				}
				tmp.isSpace = true
				tmp.id = 0
				
				break
			}
		}
		tmp = tmp.prev
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
