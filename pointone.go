package main

import "piscine"

import "fmt"

func PointOne(n *int) {
	*n = 1
}

func main() {
	n := 0
	piscine.PointOne(&n)
	
	fmt.Println(n)	
}
