package main

import "fmt"

func main() {
	age := 32

	fmt.Println("Age:", age)

	getAdultYears(&age)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
