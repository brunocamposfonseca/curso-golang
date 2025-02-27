package main 

import ("fmt")

func main(){
	//Variável Global string
	var name string = "Bruno"
	fmt.Println(name)
	name = "Josisvaldo"
	
	//Variável Global int
	var age int = 17
	fmt.Println(name,"has",age,"years old")
	
	// Constante int
	const cpf int = 11111111122
	fmt.Println(name+", CPF owner: ",cpf)
	
	var status bool = true
	fmt.Println(status)
	
	//Variável Bloco string
	fruit := "apple"
	fmt.Println(fruit)
}