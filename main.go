package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDoRenato := contas.ContaCorrente{Titular: "Renato", NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.5}
	contaDoRenato2 := contas.ContaCorrente{Titular: "Renato", NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.5}
	fmt.Println(contaDoRenato == contaDoRenato2) //true

	contaDaBruna := contas.ContaCorrente{"Bruna", 222, 112223, 200}
	contaDoGuilherme := contas.ContaCorrente{Titular: "Guilherme", Saldo: 150.5}

	fmt.Println(contaDoRenato)
	fmt.Println(contaDaBruna)
	fmt.Println(contaDoGuilherme)

	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular = "Cris"
	contaDaCris.Saldo = 500

	var contaDaCris2 *contas.ContaCorrente
	contaDaCris2 = new(contas.ContaCorrente)
	contaDaCris2.Titular = "Cris"
	contaDaCris2.Saldo = 500
	fmt.Println(contaDaCris == contaDaCris2)   //false
	fmt.Println(*contaDaCris == *contaDaCris2) //true

	fmt.Println(*contaDaCris)

	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 500.

	fmt.Println(contaDaSilvia.Saldo)
	fmt.Println(contaDaSilvia.Sacar(200))
	fmt.Println(contaDaSilvia.Saldo)

	status, valor := contaDaSilvia.Depositar(1000)
	fmt.Println(status, valor)

	resultado := contaDaSilvia.Transferir(100, contaDaCris)
	fmt.Print(resultado, contaDaSilvia.Saldo, contaDaCris.Saldo)
}
