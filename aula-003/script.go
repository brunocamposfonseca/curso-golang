package main

import (
	"fmt"
)

func main(){
	//%T = Exibe o tipo da variável; %v = Exibe o valor da variável
	fmt.Printf("Type: %T - Value: %v\n", true, true)
	fmt.Printf("Type: %T - Value: %v\n", "steph", "steph")
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)
	fmt.Printf("Type: %T - Value: %v\n", "1", "1")
	fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.233)
}