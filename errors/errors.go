package errors

import (
	`errors`
	`fmt`

	`github.com/k0kubun/pp`
)

// Работа с ошибками
type User struct {
	Name    string
	Balance int
}

const (
	noMoney = "Недостаточно средств"
)

func checkOperation(startBalance, paymentAmount int) {
	ballance := startBalance
	payment := paymentAmount

	user := User{
		Name:    "Пользователь",
		Balance: ballance,
	}

	pp.Println(
		"Баланс до произведения оплаты: ",
		user.Balance,
	)

	err := pay(
		&user,
		payment,
	)

	if err != nil {
		fmt.Println(
			"Произошла ошибка оплаты: ",
			err.Error(),
		)
	}

	pp.Println(
		"Баланс после произведения оплаты: ",
		user.Balance,
	)
}

func pay(
	user *User,
	usd int,
) error {
	// Проверяем, хватает ли на балансе пользователя суммы для соверщения покупки
	if (user.Balance - usd) < 0 {
		return errors.New(noMoney)
	}

	// Если хватает, то проводим оплату, списываем деньги
	user.Balance -= usd

	// Возращаем nil
	return nil
}
