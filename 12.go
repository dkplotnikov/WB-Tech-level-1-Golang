package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

func main() {

	list := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})
	for _, e := range list {
		set[e] = struct{}{}
	}

	for e := range set {
		fmt.Print(e, " ")
	}
	fmt.Println()
}
