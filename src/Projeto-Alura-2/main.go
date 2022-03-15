package main

import (
	"Golang-Studies/src/Projeto-Alura-2/Contas"
	"fmt"
)

func main() {
	//esse primeiro não é muito recomendado
	contaGuilherme := conta.ContaCorrente{"Guilherme", 529, 123456, 125.50}
	contaFabio := conta.ContaCorrente{
		Titular:       "Fabio",
		NumeroAgencia: 528,
		NumeroConta:   123457,
		Saldo:         0,
	}
	var contaJulia *conta.ContaCorrente = new(conta.ContaCorrente)
	contaJulia.Titular = "Julia"

	fmt.Println(contaGuilherme)
	fmt.Println(contaFabio)
	fmt.Println(*contaJulia)

	fmt.Println(contaGuilherme.Sacar(-1020))
	fmt.Println(contaGuilherme.Depositar(-1020))
	fmt.Println(contaGuilherme.Depositar(500))

	fmt.Println(contaGuilherme.Transferir(-100, &contaFabio))
	fmt.Println(contaGuilherme)
	fmt.Println(contaFabio)

}
