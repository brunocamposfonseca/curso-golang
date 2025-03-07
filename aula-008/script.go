package main

import "fmt"

func main() {
	// Map - Homogênio
	idade := map[string]int{}
	idade["João"] = 20
	idade["Maria"] = 25
	idade["José"] = 30

	fmt.Println(idade)
	fmt.Println(idade["João"])

	anoNasc := map[string]int{
		"João":  2000,
		"Maria": 1995,
		"José":  1990,
	}

	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Maria"])
	anoNasc["Ana"] = 1996
	fmt.Println(anoNasc)
}
