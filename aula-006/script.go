package main

import "fmt"

func main(){
	var x,y float64
	
	fmt.Print("Enter two numbers: ")
	fmt.Scanln(&x, &y)
	
	soma := soma(x, y)
	fmt.Println("Resultado da soma:",soma)
	
	sub := sub(x, y)
	fmt.Println("Resultado da subtração",sub)
	
	var fname string
	
	fmt.Println("Enter your first name: ")
	fmt.Scan(&fname)
	
	name, lastname := cname(fname)
	
	fmt.Println(name,lastname)
}

func cname(name string) (string, string){
	lname := "Fonseca"
	firstname := "Bruno"
	return firstname,lname
}

func soma(a float64, b float64) float64{
	return a + b 
}

func sub(a float64, b float64) float64{
	return a - b 
}