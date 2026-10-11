// В этом коде описана логика работы корзины,
// в которую можно добавлять продукты, убирать их, а также создавать заказ.

// Нужно провести Code Review и предложить улучшения по данному функционалу.

package main

import (
	"fmt"
	"net/smtp"
)

var items []Product
var vatRate = 10.0
var removalChannel chan bool
var order Order

type Product struct {
	NameOfProduct  string
	PriceOfProduct float32
}

type Order struct {
	Items []Product
}

func init() {
	items = make([]Product, 100)
	removalChannel = make(chan bool, 1)
}

func AddItemToShoppingCart(item Product) {
	items = append(items, item)
}

func RemoveItemFromShoppingCart(nameOfItem string) {
	go func() {
		for i, item := range items {
			if item.NameOfProduct == nameOfItem {
				items = append(items[:i], items[i+1:]...)
				removalChannel <- true

				return
			}
		}

		removalChannel <- false
	}()

	<-removalChannel
}

func CalculateVAT() float64 {
	subTotal := 0.0
	for _, item := range items {
		price := float64(item.PriceOfProduct)
		subTotal += price
	}

	vatAmount := subTotal * (vatRate / 100)

	return vatAmount
}

func ListAndCalculateVAT() float64 {
	subTotal := 0.0
	for _, item := range items {
		fmt.Printf("Item Name: %s, Item Price: $%.2f\n", item.NameOfProduct, item.PriceOfProduct)
		subTotal += float64(item.PriceOfProduct)
	}

	vatAmount := CalculateVAT()
	totalWithVAT := subTotal + vatAmount

	fmt.Printf("The total price with VAT: $%.2f\n", totalWithVAT)

	return totalWithVAT
}

func CreateOrder(customerEmail string) {
	totalPrice := ListAndCalculateVAT()

	orderDetails := "Thank you for your order!\n\n"
	orderDetails += "Here are the items in your order:\n"

	for _, item := range items {
		orderDetails += fmt.Sprintf("- %s: $%.2f\n", item.NameOfProduct, item.PriceOfProduct)
	}

	orderDetails += fmt.Sprintf("\nTotal (including VAT): $%.2f\n", totalPrice)

	subject := "Your Order Confirmation"

	err := SendOrderConfirmationEmail(customerEmail, subject, orderDetails)

	order = Order{Items: items}

	if err != nil {
		fmt.Println("Failed to send order confirmation email:", err)
	}
}

func SendOrderConfirmationEmail(to string, subject string, body string) error {
	from := "your-email@example.com"
	password := "your-email-password"
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	message := []byte(
		"Subject: " + subject + "\r\n" +
			"To: " + to + "\r\n" +
			"From: " + from + "\r\n" +
			"\r\n" + body,
	)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		return err
	}

	return nil
}

/*
Баги:
1. items = make([]Product, 100), баг, инициализация корзины на 100 пустых элементов

Глобальные переменные:
1. items - глобальная переменная, к которой имеют доступ все юзеры - необходима изоляция
2. Order - глобальная переменная, каждый пользователь перезаписывает
3. ставка НДС vatRate - глобальная переменная
Ставка может различаться по стране, категории товара и типу операции.
Лучше передавать налоговую ставку через конфигурацию или отдельную зависимость.
4. Ненужный канал removalChannel

AddItemToShoppingCart
1. data race при работе со слайсом items в AddItemToShoppingCart (нужно добавить mutex)

RemoveItemFromShoppingCart
1. ненужная горутина
Фактически это синхронная операция, искусственно усложнённая конкурентностью
2 Глобальный канал removalChannel может приводить к логическим ошибкам
Одна функция может получить результат другой операции.
Пункты 1 и 2 ненужный overengineering
3. При удалении RemoveItemFromShoppingCart поиск идет по имени NameOfProduct
Если два товара имеют одинаковое имя, удалится только первый
При удалении необходимо использовать уникальный id
4. Для хранения Items лучше использовать мапу с ключом в виде idProduct -
уникальность, константная запись, поиск и удаление

CalculateVAT
1. float32 — плохой выбор для денежных значений, т.к. имеет приблизительную точность
Для хранения денег используют:
- int64 - сумму в минимальных денежных единицах (копейках или центах)
- специализированный тип денег или библиотеку, когда нужна более сложная арифметика

ListAndCalculateVAT
1. Расчет суммы и вывод в 1 функции ListAndCalculateVAT,
нарушение принципа single responsibility

CreateOrder
1. CreateOrder смешивает слишком много ответственностей и
нарушает принцип single responsibility
Логичнее разделить работу на отдельные компоненты:
- CartService - управление позициями корзины
- OrderService - создание и сохранение заказа
- EmailSender - отправка уведомления
- OrderRepository - хранение заказов
Не обязательно создавать отдельный сервис для каждой мелочи,
но бизнес-логику необходимо разделять хотя бы на отдельные функции
2. Ошибки при расчете VAT
сумма товаров subTotal считается 2 раза
CalculateVAT - считает только ндс, а потом прибавляет к повторно сложенному subtotal
3. Заказ - никуда не сохраняется, никак не обрабатывается, никуда не пушится,
просто перезаписывает глобальную переменную order
Заказ и корзина ссылаются на один слайс Items, который может измениться после оформления заказа
- Заказ необходимо записывать в бд, пушить в очередь + логировать сообщение (outbox)
для идемпотентной обработки

SendOrderConfirmationEmail
1. доступы захардкожены
- пароли, порты должны быть вынесены в файлы зависимостей (env, yaml) либо хранилище паролей
- юзер должен передаватьтся аргументах функции либо лежать в middleware
2. message - собирается через конкатенацию
- сборка должна идти через strings.Builder либо использоваться готовый шаблон для писем
3. Заголовки должны корректно обрабатываться, чтобы пользовательский ввод не мог внедрить дополнительные заголовки.
Текст сообщения и параметры должны валидироваться
4. функция напрямую зависит от SMTP smtp.SendMail(...)
Это делает отправку письма жёстко привязанной к конкретной реализации.
В Go распространённый подход — зависимость от небольшого интерфейса:
type EmailSender interface {
    Send(to, subject, body string) error
}
Отправитель писем - это должен быть переиспользуемый синглтон, 1 на весь сервис
Тогда бизнес-логика зависит от контракта, а SMTP-реализация передаётся извне.
В тестах можно использовать mock или fake.
*/
