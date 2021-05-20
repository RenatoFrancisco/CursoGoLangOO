package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	contaDoRenato := contas.ContaCorrente{Titular: clientes.Cliente{Nome: "Renato"}, NumeroAgencia: 589, NumeroConta: 123456}
	contaDoRenato2 := contas.ContaCorrente{Titular: clientes.Cliente{Nome: "Renato"}, NumeroAgencia: 589, NumeroConta: 123456}
	fmt.Println(contaDoRenato == contaDoRenato2) //true

	contaDaBruna := contas.ContaCorrente{Titular: clientes.Cliente{Nome: "Bruna"}, NumeroAgencia: 222, NumeroConta: 112223}
	contaDoGuilherme := contas.ContaCorrente{Titular: clientes.Cliente{Nome: "Guilherme"}}

	fmt.Println(contaDoRenato)
	fmt.Println(contaDaBruna)
	fmt.Println(contaDoGuilherme)

	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular.Nome = "Cris"
	// contaDaCris.Saldo = 500
	contaDaCris.Depositar(500)

	var contaDaCris2 *contas.ContaCorrente
	contaDaCris2 = new(contas.ContaCorrente)
	contaDaCris2.Titular.Nome = "Cris"
	// contaDaCris2.Saldo = 500
	contaDaCris2.Depositar(500)
	fmt.Println(contaDaCris == contaDaCris2)   //false
	fmt.Println(*contaDaCris == *contaDaCris2) //true

	fmt.Println(*contaDaCris)

	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular.Nome = "Silvia"
	// contaDaSilvia.Saldo = 500.
	contaDaSilvia.Depositar(500)

	fmt.Println(contaDaSilvia.ObterSaldo())
	fmt.Println(contaDaSilvia.Sacar(200))
	fmt.Println(contaDaSilvia.ObterSaldo())

	status, valor := contaDaSilvia.Depositar(1000)
	fmt.Println(status, valor)

	resultado := contaDaSilvia.Transferir(100, contaDaCris)
	fmt.Print(resultado, contaDaSilvia.ObterSaldo(), contaDaCris.ObterSaldo())
}
