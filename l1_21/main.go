package main

import "fmt"

/*
Основная применимость паттерна Адаптер заключается в интеграции старых
или сторонних компонентов в новую систему без изменения их исходного кода.
Это обеспечивает гибкость и повторное использование, поскольку позволяет использовать
уже проверенные и надёжные классы, даже если их API не совпадают.

Основной плюс паттерна адаптер - сокрытие от клиента преобразование интерфейсов
Основной минус - усложнение кодовой базы

Пример
Если есть старая система, которая работает с базой данных через класс LegacyDatabaseConnector
с методом queryData(String sql). А команда разрабатывает новую систему, которая использует
современный интерфейс IDataProvider с методом fetchRecords(String query).
Вместо того чтобы переписывать всю старую логику, можно создать класс-адаптер
*/

// Producer Внешний пакет недоступный для редактирования
/*******************************************************/

type ProducerI interface {
	Produce()
}

type Producer struct {
}

func (c *Producer) Produce() {
	fmt.Println("Produce")
}

/*******************************************************/

// Consumer собственный пакет
/*******************************************************/

type ConsumerI interface {
	Consume()
}

// Адаптер

type ProducerAdapter struct {
	producer *Producer
}

func (a *ProducerAdapter) Consume() {
	fmt.Println("Adapter Consume")
	a.producer.Produce()
}

/*******************************************************/

func main() {
	producer := &Producer{}
	var consumerClient ConsumerI = &ProducerAdapter{producer: producer}

	consumerClient.Consume()
}
