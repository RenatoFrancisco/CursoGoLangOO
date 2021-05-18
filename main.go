package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoRenato := ContaCorrente{titular: "Renato", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	contaDaBruna := ContaCorrente{"Bruna", 222, 112223, 200}
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 150.5}

	fmt.Println(contaDoRenato)
	fmt.Println(contaDaBruna)
	fmt.Println(contaDoGuilherme)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500
	fmt.Println(*contaDaCris)
}
