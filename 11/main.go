package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func findSetIntersections(set1, set2 map[string]struct{}) []string {
	var set []string
	hash := make(map[string]struct{})
	for i, v := range set1 {
		hash[i] = v
	}

	for i := range set2 {
		if _, ok := hash[i]; ok {
			set = append(set, i)
		}
	}
	return set
}

func main() {
	set1 := map[string]struct{}{
		"apple":  {},
		"lemon":  {},
		"orange": {},
		"grapes": {},
	}
	set2 := map[string]struct{}{
		"apple":  {},
		"lemon":  {},
		"orange": {},
		"banana": {},
		"peach":  {},
	}
	fmt.Println(findSetIntersections(set1, set2))
}
