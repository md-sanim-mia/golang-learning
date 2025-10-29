package main

import "fmt"

type payment struct {
	getwa stripe
}

func (p payment) makePament(amount float32) {

	p.getwa.pay(amount)

}

type stripe struct{}

func (s stripe) pay(amount float32) {

	fmt.Println("make payment with stripe", amount)
}
func main() {

	newPayment := payment{}

	newPayment.makePament(45455)
}
