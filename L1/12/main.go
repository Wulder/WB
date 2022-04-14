package main

import (
	"fmt"
)

func main() {

	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	multiplicity := make(map[string]struct{}) //множество

	for _, s := range strings {
		multiplicity[s] = struct{}{}
	}

	fmt.Println(multiplicity)
}
