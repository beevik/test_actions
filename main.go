package main

import (
	"fmt"

	"github.com/beevik/prefixtree/v2"
)

func main() {
	prefixtree_test()
}

func prefixtree_test() {
	tree := prefixtree.New[int]()

	tree.Add("apple", 10)
	tree.Add("orange", 20)
	tree.Add("apple pie", 30)
	tree.Add("lemon", 40)
	tree.Add("lemon meringue pie", 50)

	fmt.Printf("%-18s %-8s %s\n", "prefix", "value", "error")
	fmt.Printf("%-18s %-8s %s\n", "------", "-----", "-----")

	for _, prefix := range []string{
		"a",
		"appl",
		"apple",
		"apple p",
		"apple pie",
		"apple pies",
		"o",
		"oran",
		"orange",
		"oranges",
		"l",
		"lemo",
		"lemon",
		"lemon m",
		"lemon meringue",
		"pear",
	} {
		value, err := tree.FindValue(prefix)
		fmt.Printf("%-18s %-8v %v\n", prefix, value, err)
	}
}
