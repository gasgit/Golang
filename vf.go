package main

import "fmt"

// (args ...int) zero or more args

// create function to add n args
func addition(args ...int) int{
	// var to sum args
	total := 0
	for _, val := range args{
		total += val
	}
	return total
}


func main(){

	fmt.Println(addition(2,5,7))
	fmt.Println(addition(1,2,5,6,7))

}

