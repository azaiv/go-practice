package errors

import `fmt`

// Обработка panic, что бы приложение не падало
func checkPanic() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println(
				"Отработал panic",
				p,
			)
		}
	}()

	slice := []int{
		1,
		2,
		3,
	}

	fmt.Println(slice[4])
}
