package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 24,
		"Unborn":          0,
	}
	fmt.Println("map:\t\t", m)
	fmt.Println("m[\"James\"]:\t", m["James"])
	fmt.Println("m[\"Unborn\"]:\t", m["Unborn"])

	fmt.Println("\nLook at what happens when you use a key that does not exist:")
	fmt.Println("m[\"Nilson\"]:\t", m["Nilson"])

	fmt.Println("\nFortunately, using a key on a map returns two values:")
	v, ok := m["Nilson"]
	fmt.Println("v, ok := m[\"Nilson\"]\t\tv =", v, "\tok =", ok)
	v1, ok1 := m["Unborn"]
	fmt.Println("v, ok := m[\"Unborn\"]\t\tv =", v1, "\tok =", ok1)
	v2, ok2 := m["James"]
	fmt.Println("v, ok := m[\"James\"]\t\tv =", v2, "\tok =", ok2)
	v3, ok3 := m["Miss Moneypenny"]
	fmt.Println("v, ok := m[\"Miss Moneypenny\"]\tv =", v3, "\tok =", ok3)

	// ADD TO A MAP
	fmt.Println("\n---------------")
	fmt.Println("- ADD TO MAP  -")
	fmt.Println("---------------")

	fmt.Println("To add to a map:\t m[\"key\"] = \"value\"")
	fmt.Println("map:\t\t\t", m)

	m["Nilson"] = 26
	fmt.Println("\nPerform add:\t\t m[\"Nilson\"] = 26")
	fmt.Println("map:\t\t\t", m)

	// LOOP THROUGH A MAP
	fmt.Println("\n---------------")
	fmt.Println("-   ITERATE   -")
	fmt.Println("---------------")

	for k, v := range m {
		fmt.Println(k, v)
	}

	// DELETE FROM A MAP
	fmt.Println("\n---------------")
	fmt.Println("-   DELETE    -")
	fmt.Println("---------------")

	fmt.Println("To delete from a map:\t delete(<map name>, \"key\")")
	fmt.Println("map:\t\t\t", m)

	fmt.Println("\nPerform delete:\t\t delete(m, \"Unborn\")")
	delete(m, "Unborn")
	fmt.Println("map:\t\t\t", m)

	fmt.Println("\nDelete non-existent:\t delete(m, \"Unborn\")")
	delete(m, "Unborn")
	fmt.Println("map:\t\t\t", m)

	fmt.Println("\nTo be sure you delete, use the 'comma ok idiom'")
	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:\t", v)
		delete(m, "Miss Moneypenny")
	}
	fmt.Println("map:\t", m)

}
