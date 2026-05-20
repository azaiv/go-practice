package methods

import (
	"fmt"
	"math/rand"
)

type Paypal struct{}

func NewPaypal() Paypal {
	return Paypal{}
}

func (c Paypal) Pay(usd int) int {
	fmt.Println("Оплата через Paypal")
	fmt.Println("Размер опалты: ", usd, "USD")

	return rand.Int()
}

func (c Paypal) Cancel(id int) {
	fmt.Println("Отмена Paypal операции ID:", id)
}
