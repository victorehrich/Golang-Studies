package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaGuilherme := ContaCorrente{"Guilherme", 529, 123456, 125.50}
	contaFabio := ContaCorrente{
		titular:       "Fabio",
		numeroAgencia: 528,
		numeroConta:   123457,
		saldo:         0,
	}
	var contaJulia *ContaCorrente = new(ContaCorrente)
	contaJulia.titular = "Julia"

	fmt.Println(contaGuilherme)
	fmt.Println(contaFabio)
	fmt.Println(*contaJulia)

	fmt.Println(contaGuilherme.Sacar(-1020))
	fmt.Println(contaGuilherme)

}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque > 0 {
		if c.saldo > valorDoSaque {
			c.saldo -= valorDoSaque
			return "Saque realizado com sucesso"
		}
		return "Saldo insuficiente"
	}
	return "Valor Inválido"

}
