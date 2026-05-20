package payment

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentInfo   map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentInfo:   make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

// Метод Pay()
// Принимает:
// 1. Описание проводимой оплаты
// 2. Сумма оплаты
// Возращает:
// ID проводимой операции

func (p PaymentModule) Pay(
	description string,
	usd int,
) int {
	// 1. Проводить оплату
	// 2. Получать ID проведенной операции
	id := p.paymentMethod.Pay(usd)

	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Cancelled:   false,
	}

	// 3. Сохранять информацию о проведенной операции
	// - Описание операции
	// - Сколько было потрачено
	// - Отмененная операция
	p.paymentInfo[id] = info

	// 4. Возвращать ID проведенной операции
	return id
}

// Метод Cancel()
// Принимает:
// 1. ID операции
// Возращает:
// - ничего не возращает

func (p PaymentModule) Cancel(id int) {
	info, ok := p.paymentInfo[id]

	if !ok {
		return
	}

	p.paymentMethod.Cancel(id)

	info.Cancelled = true

	p.paymentInfo[id] = info
}

// Метод Info()
// Принимает:
// 1. ID операции
// Возращает:
// Информацию о проведенной операции

func (p PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentInfo[id]

	if !ok {
		return PaymentInfo{}
	}

	return info
}

// Метод Info()
// Принимает:
// -
// Возращает:
// О всех проведенных операциях

func (p *PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(
		map[int]PaymentInfo,
		len(p.paymentInfo),
	)

	for k, v := range p.paymentInfo {
		tempMap[k] = v
	}

	return tempMap
}
