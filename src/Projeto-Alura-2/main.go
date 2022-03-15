package main

import (
	"Golang-Studies/src/Projeto-Alura-2/Clientes"
	"Golang-Studies/src/Projeto-Alura-2/Contas"
	"fmt"
)

func main() {
	contaBruno := conta.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Bruno",
			CPF:       "123.123.123.12",
			Profissao: "Dev",
		},
		NumeroAgencia: 123,
		NumeroConta:   123123,
	}

	clienteAlysson := clientes.Titular{
		Nome:      "Alysson",
		CPF:       "123.123.123.13",
		Profissao: "PO",
	}
	contaAlysson := conta.ContaCorrente{
		Titular:       clienteAlysson,
		NumeroAgencia: 123,
		NumeroConta:   123124,
	}
	fmt.Println(contaBruno)
	fmt.Println(contaAlysson)

}
