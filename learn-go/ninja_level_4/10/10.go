package main

import "fmt"

/* HANDS ON # 10 -----------------------------------
1. using the code from the previous example, delete a record to your map.
2. now print the map out using the “range” loop.
----------------------------------- */
func main() {
	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	m["mcleod_todd"] = []string{"Surfing", "Learning", "Teaching"}
	delete(m, "no_dr")

	for k, v := range m {
		fmt.Printf("%v:\n", k)
		for i, thing := range v {
			fmt.Println(" ", i, ":", thing)
		}
	}
}
