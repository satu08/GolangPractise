package main

import (
	"encoding/json"
	"fmt"
)

type Male struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"satya","Last":"jadhav","Age":23},{"First":"saurabh","Last":"vidhate","Age":24}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
	var male []Male
	err := json.Unmarshal(bs, &male)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(male)
	for i, v := range male {
		fmt.Println(i)
		fmt.Println(v.First, v.Last, v.Age)
	}
	//json.Unmarshal(make([]byte, 0), &male)
}
