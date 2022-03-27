package conta

import clientes "Programs/Golang-Studies/src/Projeto-Alura-2/Clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) MostrarSaldo() float64 {
	return c.saldo
}
func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	if valorDoSaque > 0 {
		if c.saldo > valorDoSaque {
			c.saldo -= valorDoSaque
			return "Saque realizado com sucesso", c.saldo
		}
		return "saldo insuficiente", -1
	}
	return "Valor Inválido", -1
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso", c.saldo
	}
	return "Valor Inválido", -1
}

func (cs *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) string {
	if valorTransferencia > 0 && cs.saldo > valorTransferencia {
		cs.Sacar(valorTransferencia)
		contaDestino.Depositar(valorTransferencia)
		return "Transferencia realizada com sucesso"
	}
	return "Valor Inválido"
}
