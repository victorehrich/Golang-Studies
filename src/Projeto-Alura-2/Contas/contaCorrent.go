package conta

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	if valorDoSaque > 0 {
		if c.Saldo > valorDoSaque {
			c.Saldo -= valorDoSaque
			return "Saque realizado com sucesso", c.Saldo
		}
		return "Saldo insuficiente", -1
	}
	return "Valor Inválido", -1
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.Saldo += valorDeposito
		return "Deposito realizado com sucesso", c.Saldo
	}
	return "Valor Inválido", -1
}

func (cs *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) string {
	if valorTransferencia > 0 && cs.Saldo > valorTransferencia {
		cs.Sacar(valorTransferencia)
		contaDestino.Depositar(valorTransferencia)
		return "Transferencia realizada com sucesso"
	}
	return "Valor Inválido"
}
