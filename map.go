package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["One"] = 1
	x["Two"] = 2
	x["Three"] = 3

	fmt.Println(x, len(x))
	fmt.Println(x["Three"])

	delete(x, "One")

	fmt.Println(x)

	// check for element
	if item, ok := x["Two"]; ok {
		fmt.Println(item, ok)
	}

	chemicalElements()

	nestedMap()
}

func chemicalElements() {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
	}

	fmt.Println(elements)
}

func nestedMap() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}

	fmt.Println(elements["Li"])

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
