package goroutines

import (
	"fmt"
	"time"
)

func Run() {
	coal := 0

	initTime := time.Now()

	// Каналы можно передавать копией
	transferPoint := make(chan int, 2)

	for i := 1; i <= 3; i++ {
		go mine(
			transferPoint,
			i,
		)
		coal += <-transferPoint
	}

	fmt.Println(
		"Добыли ",
		coal,
		" количеств угля",
	)
	fmt.Println(
		"Прошло времени",
		time.Since(initTime),
	)
}
