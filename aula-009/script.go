package main

import "fmt"

type Pessoa struct{
	Id int
	Nome string
	Sobrenome string
	Idade int
	Cpf string
}

type Aluno struct {
	Pessoa
	Curso string
}

func main(){
	//Structs
	p1 := Pessoa{1, "João", "Silva", 30, "123.456.789-00"}
	p2 := Pessoa{2, "Maria", "Santos", 25, "987.654.321-00"}

	fmt.Println(p1)
	fmt.Println(p2)
	
	fmt.Println(Pessoa{Id: 3, Nome: "José", Sobrenome: "Oliveira", Idade: 40, Cpf: "456.789.123-00"})
	
	p3 := Pessoa{Id: 4, Nome: "Ana", Sobrenome: "Souza", Idade: 35, Cpf: "789.123.456-00"}
	fmt.Println(p3)

	p1.Cpf = ""

	fmt.Println(NomeCompleto(p1.Nome, p1.Sobrenome) + ", teve o CPF desativado")

	pList := []Pessoa{}

	pList = append(pList, p1, p2, p3)
	fmt.Println(pList)

	alunos := map[string][]Pessoa{}
	alunos["ds"] = pList
	fmt.Println(alunos["ds"])
	alunos["multimidia"] = []Pessoa{p1}
	fmt.Println(alunos["multimidia"])

	//Herança

	aluno := Aluno{p1, "Engenharia de Software"}
	fmt.Println(aluno)
	fmt.Println(aluno.Pessoa.Nome)
	fmt.Println(aluno.Sobrenome)
	fmt.Println(aluno.Curso)

}

func NomeCompleto(name string, lastname string) string {	
	return name + " " + lastname
}