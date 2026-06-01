package goroutines

import (
	"fmt"
	"time"
)

// Функция, описывающая поход рабочего в шахту
// Принимает:
// 1. Номер похода в шахту
// Возращает
// 1. Добытый уголь
func mine(
	transferPoint chan int,
	n int,
) {
	fmt.Println(
		"Поход в шахту номер ",
		n,
		"начался",
	)

	time.Sleep(1 * time.Second)

	fmt.Println(
		"Поход в шахту номер ",
		n,
		"закончился",
	)

	transferPoint <- 10
}
