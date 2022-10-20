package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, word := range words {
		set[word] = struct{}{}
	}

	fmt.Println(set)
}
