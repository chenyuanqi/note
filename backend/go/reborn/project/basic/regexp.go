package main

import (
	"fmt"
	"regexp"
)

// https://yourbasic.org/golang/regexp-cheat-sheet/
func main() {
	// check if there is a substring matching a.b
	matched, err := regexp.MatchString(`a.b`, "aaxbb")
	fmt.Println(matched) // true
	fmt.Println(err)     // nil (regexp is valid)
}
