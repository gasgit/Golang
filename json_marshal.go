package main

import (
	"encoding/json"
	"fmt"
)

// create struct for movie
type movie struct {
	Title string
	Year  int
	Rate  float32
	Cast  []string
}

func main() {

	// craete movies
	m1 := movie{
		Title: "The Shawshank Redemption",
		Year:  1994,
		Rate:  9.4,
		Cast: []string{
			"Tim Robbins",
			"Morgan Freeman",
			"Bob Gunton",
		},
	}
	m2 := movie{

		Title: "The Godfather",
		Year:  1972,
		Rate:  9.2,
		Cast: []string{
			"Marlon Brando",
			"Al Pacino",
			"James Caan",
		},
	}

	fmt.Println(m1)

	// create slice movie
	movie_collection := []movie{m1, m2}

	// marshal movie_collection to byte stream
	//bStream, err := json.Marshal(movie_collection)

	// go pretty print
	bStream, err := json.MarshalIndent(movie_collection, "", " ")
	// error check
	if err != nil {
		fmt.Println(err)
	}
	// convert bs to string and print
	fmt.Println(string(bStream))

}
