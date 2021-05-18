package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Sque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaDoRenato := ContaCorrente{titular: "Renato", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	contaDoRenato2 := ContaCorrente{titular: "Renato", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	fmt.Println(contaDoRenato == contaDoRenato2) //true

	contaDaBruna := ContaCorrente{"Bruna", 222, 112223, 200}
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 150.5}

	fmt.Println(contaDoRenato)
	fmt.Println(contaDaBruna)
	fmt.Println(contaDoGuilherme)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"
	contaDaCris2.saldo = 500
	fmt.Println(contaDaCris == contaDaCris2)   //false
	fmt.Println(*contaDaCris == *contaDaCris2) //true

	fmt.Println(*contaDaCris)

	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500.

	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.Sacar(200))
	fmt.Println(contaDaSilvia.saldo)
}
