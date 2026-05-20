package errors

func RunErrors(startBalance, paymentAmount int) {
	checkOperation(
		startBalance,
		paymentAmount,
	)
}

func RunPanic() {
	checkPanic()
}
