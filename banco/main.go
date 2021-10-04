package main

import (
	"alura/banco/contas"
	"fmt"
)

//Função da interface
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

//Implementar interface
type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	//Interface
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 1000)

	fmt.Println(contaDaLuisa.ObterSaldo())

	//=================================================================================================================
	//Classes no mesmo pacote
	//contaDoDenis := contas.ContaPoupanca{}
	//contaDoDenis.Depositar(100)

	//fmt.Println(contaDoDenis.ObterSaldo())

	//=================================================================================================================

	//clienteBruno := cl.Titular{Nome: "Bruno", CPF: "123.123.123-99", Profissao: "Desenvolvedor"}
	//contaDoBruno := c.ContaCorrente{clienteBruno, 123, 123456, 100}
	//fmt.Println(contaDoBruno)

	//=================================================================================================================

	//contaDaSilvia := c.ContaCorrente{Titular: "Silvia", Saldo: 300}
	//contaDoGustavo := c.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	//status := contaDaSilvia.Transferir(150, &contaDoGustavo)

	//fmt.Println(status)
	//fmt.Println(contaDaSilvia)
	//fmt.Println(contaDoGustavo)

	//=================================================================================================================

	//contaDaSilvia := ContaCorrente{}
	//contaDaSilvia.titular = "Silvia"
	//contaDaSilvia.saldo = 500

	//valorDoSaque := 200. //Usa-se o . para ele etender que é do tipo float
	//contaDaSilvia.saldo = contaDaSilvia - valorDoSaque

	//status, valor := contaDaSilvia.Depositar(100)
	//fmt.Println(contaDaSilvia.Sacar(200))
	//fmt.Println(contaDaSilvia.saldo)
	//fmt.Println(status)
	//fmt.Println(valor)
	//fmt.Println(contaDaSilvia.saldo)

	//=================================================================================================================

	//contaDoGuilherme := ContaCorrente{titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 123.5}
	//	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	//	fmt.Println(contaDoGuilherme)
	//	fmt.Println(contaDaBruna)
	//
	//	var contaDaCris *ContaCorrente //O * indica o ponteiro de memoria (nesse caso se tentar compara com outra variavel de mesmo valor vai dar falso)
	//	contaDaCris = new(ContaCorrente)
	//	contaDaCris.titular = "Cris"
	//	contaDaCris.saldo = 500
	//
	//	fmt.Println(*contaDaCris)
	//	fmt.Println(&contaDaCris) //& retorna o endereço de memoria
}
