package main

// Human Структура
type Human struct {
	// Поля структуры
	Name    string
	Surname string
	Address string
}

// Методы структуры Human
func (h *Human) GetFullName() string {
	return h.Name + " " + h.Surname
}

func (h *Human) GetAddress() string {
	return h.Address
}

// Action Структура, которая использует композицию с Human
type Action struct {
	Human
}

// Методы структуры Action, которые используют методы структуры Human
func (a *Action) makeOrder() string {
	return "Order made by " + a.GetFullName() + " at address: " + a.GetAddress()
}

func (a *Action) cancelOrder() string {
	return "Order cancelled by " + a.GetFullName() + " at address: " + a.GetAddress()
}

func main() {

	// Создание экземпляра структуры Action с вложенной структурой Human
	action := Action{
		Human: Human{
			Name:    "John",
			Surname: "Doe",
			Address: "123 Main St",
		},
	}

	// Использование методов структуры Action
	// которые в свою очередь используют методы структуры Human
	println(action.GetFullName())
	println(action.GetAddress())

	// Выполнение действий с заказами
	println(action.makeOrder())
	println(action.makeOrder())
	println(action.cancelOrder())

}
