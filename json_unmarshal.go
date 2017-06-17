package main

import (
	"encoding/json"
	"fmt"
)

type movie struct {
	Title string
	Year  int
	Rate  float32
	Cast  []string
}

func main() {

	// marshaled data
	mjson := `[{"Title":"The Shawshank Redemption","Year":1994,"Rate":9.4,"Cast":["Tim Robbins","Morgan Freeman","Bob Gunton"]},{"Title":"The Godfather","Year":1972,"Rate":9.2,"Cast":["Marlon Brando","Al Pacino","James Caan"]}]`

	fmt.Println(mjson)

	// slice movies data structure
	unjson := []movie{}
	// pass slice bytes and address of slice of movie data structure
	err := json.Unmarshal([]byte(mjson), &unjson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(unjson)

	for i, v := range unjson {
		fmt.Println(i, v)
		fmt.Println(v.Title)
		fmt.Println(v.Year)
		fmt.Println(v.Cast)
		fmt.Println(v.Cast[1])

	}

}
