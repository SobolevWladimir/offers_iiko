package status

import "gopkg.in/guregu/null.v3"

//новый
var StatusNewOrder = Status{

	Name:     "Новый Заказ",
	Code:     "new",
	Color:    "#27509e",
	Comment:  null.NewString("Новый заказ", false),
	Priority: 1,
}

// Ожидает
var StatusAwaitOrder = Status{

	Name:     "Ожидает",
	Code:     "await",
	Color:    "#fad13",
	Comment:  null.NewString("Заказ в ожидании ", false),
	Priority: 2,
}

// Готовится
var StatusPrepareOrder = Status{

	Name:     "Готовится",
	Code:     "pripare",
	Color:    "#fad13",
	Comment:  null.NewString("Заказ Готовится", false),
	Priority: 3,
}

// Приготовлен
var StatusCookedOrder = Status{

	Name:     "Приготовлен",
	Code:     "cooked",
	Color:    "#3baa36",
	Comment:  null.NewString("Заказ Приготовлен", false),
	Priority: 4,
}

// Ждет отправки
var StatusWaitingSendOrder = Status{

	Name:     "Ждет отправки",
	Code:     "waitingSend",
	Color:    "#3baa36",
	Comment:  null.NewString("Заказ закрыт", false),
	Priority: 5,
}

// Отправлен
var StatusShippedOrder = Status{

	Name:     "Отправлен",
	Code:     "shipped",
	Color:    "#3baa36",
	Comment:  null.NewString("Заказ отправлен", false),
	Priority: 6,
}

// Доставлен
var StatusDeliveredOrder = Status{

	Name:     "Доставлен",
	Code:     "delivered",
	Color:    "#3baa36",
	Comment:  null.NewString("Заказ Доставлен", false),
	Priority: 7,
}

// Закрыт
var StatusClosedOrder = Status{

	Name:     "Закрыт",
	Code:     "сlosed",
	Color:    "#3baa36",
	Comment:  null.NewString("Заказ закрыт", false),
	Priority: 8,
}

// Отменет
var StatusCancelOrder = Status{

	Name:     "Отменен",
	Code:     "cancel",
	Color:    "#d72125",
	Comment:  null.NewString("Заказ отменен", false),
	Priority: 0,
}
