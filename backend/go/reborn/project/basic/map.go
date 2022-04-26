package main

import "fmt"

func dump(params ...interface{}) {
	fmt.Println(params...)
}

func main() {
	dump("Find a key in a map:")
	m := map[string]float64{"pi": 3.14}
	v, found := m["pi"] // v == 3.14  found == true
	dump(v, found)
	v, found = m["pie"] // v == 0.0   found == false
	dump(v, found)
	_, found = m["pi"] // found == true
	dump(v, found)
	// Use second return value directly in an if statement
	m = map[string]float64{"pi": 3.14}
	if v, found := m["pi"]; found {
		dump(v) // 3.14
	}
	// Check for zero value
	m = map[string]float64{"pi": 3.14}
	pi := m["pi"] // v == 3.14
	dump(pi)
	pi = m["pie"] // v == 0.0 (zero value)
	dump(pi)
}
