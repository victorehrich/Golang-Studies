package conta

import (
	clientes "Programs/Golang-Studies/src/Projeto-Alura-2/Clientes"
)

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) MostrarSaldo() float64 {
	return c.saldo
}
func (c *ContaPoupanca) Sacar(valorDoSaque float64) (string, float64) {
	if valorDoSaque > 0 {
		if c.saldo > valorDoSaque {
			c.saldo -= valorDoSaque
			return "Saque realizado com sucesso", c.saldo
		}
		return "saldo insuficiente", -1
	}
	return "Valor Inválido", -1
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso", c.saldo
	}
	return "Valor Inválido", -1
}
