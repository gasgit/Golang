package main

import "fmt"

func main() {

	m := make(map[string]int)
	fmt.Println(m)
	fmt.Println("---------------------------------------------------------------")

	m1 := map[string]int{"myKey": 11, "myKey2": 22}
	fmt.Println(m1)
	fmt.Println("---------------------------------------------------------------")

	days := map[string]int{
		"mon": 1,
		"tue": 2,
		"wed": 3,
		"thu": 4,
		"fri": 5,
		"sat": 6,
		"sun": 7,
	}

	dSlice := []string{}

	for key, _ := range days {
		dSlice = append(dSlice, key)
	}

	fmt.Println(dSlice)
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(days)
	fmt.Println("---------------------------------------------------------------")

	delete(days, "sun")
	fmt.Println(days)
	fmt.Println("---------------------------------------------------------------")

	// map of string slices
	cars := map[string][]string{
		"Ford":   {"Focus", "Mondeo"},
		"Toyota": {"Avensis", "Corolla"},
	}

	fmt.Println(cars)
	fmt.Println("---------------------------------------------------------------")

	cars["Toyota"] = append(cars["Toyota"], "Prius")

	fmt.Println(cars)
	fmt.Println("---------------------------------------------------------------")

	fmt.Println("Length Map: ", len(cars))
	fmt.Println("---------------------------------------------------------------")

	for key, value := range cars {
		fmt.Println(key, "=", value)
	}

	fmt.Println("---------------------------------------------------------------")

	if car, ok := cars["Toyota"]; ok {
		fmt.Println("Toyota", car)
	}

}
