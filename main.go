package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	contaDoRenato := contas.ContaCorrente{Titular: clientes.Cliente{Nome: "Renato"}, NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.5}
	contaDoRenato2 := contas.ContaCorrente{Titular: clientes.Cliente{Nome: "Renato"}, NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.5}
	fmt.Println(contaDoRenato == contaDoRenato2) //true

	contaDaBruna := contas.ContaCorrente{clientes.Cliente{Nome: "Bruna"}, 222, 112223, 200}
	contaDoGuilherme := contas.ContaCorrente{Titular: clientes.Cliente{Nome: "Guilherme"}, Saldo: 150.5}

	fmt.Println(contaDoRenato)
	fmt.Println(contaDaBruna)
	fmt.Println(contaDoGuilherme)

	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular.Nome = "Cris"
	contaDaCris.Saldo = 500

	var contaDaCris2 *contas.ContaCorrente
	contaDaCris2 = new(contas.ContaCorrente)
	contaDaCris2.Titular.Nome = "Cris"
	contaDaCris2.Saldo = 500
	fmt.Println(contaDaCris == contaDaCris2)   //false
	fmt.Println(*contaDaCris == *contaDaCris2) //true

	fmt.Println(*contaDaCris)

	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular.Nome = "Silvia"
	contaDaSilvia.Saldo = 500.

	fmt.Println(contaDaSilvia.Saldo)
	fmt.Println(contaDaSilvia.Sacar(200))
	fmt.Println(contaDaSilvia.Saldo)

	status, valor := contaDaSilvia.Depositar(1000)
	fmt.Println(status, valor)

	resultado := contaDaSilvia.Transferir(100, contaDaCris)
	fmt.Print(resultado, contaDaSilvia.Saldo, contaDaCris.Saldo)
}
