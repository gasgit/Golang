package main

import (
	"encoding/json"
	"fmt"
	//"os"
)

type Employee struct {
	FIRST string `json:"first_name"`
	LAST  string `json:"last_name"`
	PPS   string `json:"pps_number"`
}

type Department struct {
	DEPT  string `json:"dept_name"`
	EMAIL string `json:"email_address"`
}

type Combined struct {
	Employee
	Department
}

func main() {

	//
	emp := Employee{
		"joe",
		"blog",
		"0003",
	}

	dept := Department{
		"IT",
		"info@it.com",
	}

	com := Combined{
		emp,
		dept,
	}

	fmt.Println(com)

	bs_combined, _ := json.MarshalIndent(com, "", "  ")
	fmt.Println(string(bs_combined))
	//os.Stdout.Write(bs_combined)

	fmt.Println("------------- Unmarshal ----------------------")

	// create new data structure
	un_combined := Combined{}
	// Unmarshal bs_combined
	err := json.Unmarshal([]byte(bs_combined), &un_combined)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(un_combined)

}
