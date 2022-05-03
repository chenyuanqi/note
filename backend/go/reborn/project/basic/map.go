package main

import (
	"fmt"
	"sort"
)

func dump(params ...interface{}) {
	fmt.Println(params...)
}

func main() {
	dump("Create a new map")
	var mp map[string]int               // nil map of string-int pairs
	m1 := make(map[string]float64)      // Empty map of string-float64 pairs
	m2 := make(map[string]float64, 100) // Preallocate room for 100 entries
	m3 := map[string]float64{           // Map literal
		"e":  2.71828,
		"pi": 3.1416,
	}
	dump(mp, m1, m2, m3, len(m3)) // Size of map: 2

	dump("Add, update, get and delete keys/values")
	ma := make(map[string]float64)
	ma["pi"] = 3.14      // Add a new key-value pair
	ma["pi"] = 3.1416    // Update value
	fmt.Println(ma)      // Print map: "map[pi:3.1416]"
	v := ma["pi"]        // Get value: v == 3.1416
	v = ma["pie"]        // Not found: v == 0 (zero value)
	_, found := ma["pi"] // found == true
	_, found = ma["pie"] // found == false
	if x, found := ma["pi"]; found {
		fmt.Println(x)
	} // Prints "3.1416"
	delete(ma, "pi") // Delete a key-value pair
	fmt.Println(ma)  // Print map: "map[]"

	dump("For-each range loop")
	ml := map[string]float64{
		"pi": 3.1416,
		"e":  2.71828,
	}
	fmt.Println(ml)              // "map[e:2.71828 pi:3.1416]"
	for key, value := range ml { // Order not specified
		fmt.Println(key, value)
	}

	dump("Find a key in a map:")
	m := map[string]float64{"pi": 3.14}
	v, found = m["pi"] // v == 3.14  found == true
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

	dump("Get slices of keys and values from a map")
	keys := make([]string, 0, len(m))
	values := make([]float64, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	dump(keys, values)

	dump("Sort a map by key or value")
	mk := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25}
	mkKeys := make([]string, 0, len(mk))
	for k := range mk {
		mkKeys = append(mkKeys, k)
	}
	sort.Strings(mkKeys)
	for _, k := range mkKeys {
		fmt.Println(k, mk[k])
	}
}
