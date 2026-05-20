package _interface

import (
	`go-study/interface/payment`
	`go-study/interface/payment/methods`

	`github.com/k0kubun/pp`
)

func Run() {
	// Создаем нужный нам метод для оплаты
	method := methods.NewBank()

	// Создаем модуль для оплаты
	paymentModule := payment.NewPaymentModule(method)

	// Создаем транзацкции, храним ID операции
	idiPhone := paymentModule.Pay(
		"Покупка iPhone",
		1400,
	)
	idGooglePixel := paymentModule.Pay(
		"Покупка Google Pixel",
		1100,
	)

	allInfo := paymentModule.AllInfo()

	pp.Println(
		"Все произведенные оплаты",
		allInfo,
	)

	googleInfo := paymentModule.Info(idGooglePixel)
	iPhoneInfo := paymentModule.Info(idiPhone)

	pp.Println(googleInfo)
	pp.Println(iPhoneInfo)
}
