package main

import "fmt"

func main(){
	//Array - Tamanho fixo
	var array[2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	numPrimos := [5]int{2, 3, 5, 7, 11}
	fmt.Println(numPrimos)
	fmt.Println(len(numPrimos))

	// Slices - Tamanho ilimitado

	slice := []int{2, 3, 5, 7, 11}
	fmt.Println(slice)
	fmt.Println(slice[0:2])
	slice = append(slice, 13)
	fmt.Println(slice)
}
