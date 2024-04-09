package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	conta1 := ContaCorrente{titular: "Carlos", numeroAgencia: 01, numeroConta: 111, saldo: 0}
	conta2 := ContaCorrente{"Bruna", 01, 222, 5_000}
	conta22 := ContaCorrente{"Bruna", 01, 222, 5_000}

	fmt.Println(ContaCorrente{})
	fmt.Println(conta1)
	fmt.Println(conta2)

	var conta3 *ContaCorrente
	conta3 = new(ContaCorrente)
	// var conta3 = new(ContaCorrente)
	conta3.titular = "Pedrinho"
	fmt.Println(conta3)

	var conta33 *ContaCorrente
	conta3 = new(ContaCorrente)
	// var conta3 = new(ContaCorrente)
	conta3.titular = "Pedrinho"
	fmt.Println(conta3)
	fmt.Println(conta3 == conta33)   // false compara endereços de memoria
	fmt.Println(*conta3 == *conta33) // true

	fmt.Println(conta2 == conta22) // no go vc consegue comparar objetos true
}
