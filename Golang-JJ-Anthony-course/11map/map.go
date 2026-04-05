package main

import "fmt"

func main() {
	fmt.Println("introduction to map")

	m := map[string]int{
		"satya":   23,
		"vaibhav": 23,
		"saurabh": 24,
		"vedant":  23,
		"vishal":  24,
	}
	fmt.Println("vishal age is", m["vishal"])
	fmt.Println(m)
	fmt.Printf("%#v\n", m)
	map1 := make(map[string]int)
	map1["ravi"] = 27
	map1["swati"] = 30
	fmt.Println(map1)
	fmt.Printf("%#v\n", map1)
	fmt.Println(len(map1))

	// deleting element from map
	delete(m, "vishal")
	fmt.Println(m)
	// printing out keys and values
	for k, v := range m {
		fmt.Println(k, v)
	}
	// printing out values only
	for _, v := range m {
		fmt.Println(v)
	}
	// printing out keys only
	for k := range m {
		fmt.Println(k)
	}

	// comma ok idiom

	if v, ok := m["vishal"]; !ok {
		fmt.Println("key doesnt exist") // handle errors first
	} else {
		fmt.Println(v)
	}

	map2 := make(map[string][]string)
	map2["satya"] = []string{"satya", "siemens"}
	map2["vedant"] = []string{"vedant", "rsl"}
	map2["rahul"] = []string{"rahul", "capgemini"}
	fmt.Printf("%#v\n", map2)

	for k, v := range map2 {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}

	map2["vaibhav"] = []string{"vaibhav", "rsl"}
	fmt.Println(map2)

	delete(map2, "vedant")
	fmt.Println(map2)

}
