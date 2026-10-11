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
1. items = make([]Product, 100), баг, инициализация корзины на 100 пустых элементов
2. items - глоабальная переменная, к которой имеют доступ все юзеры - необходима изоляция
3. data race при работе со слайсом items в AddItemToShoppingCart (можно добавить mutex)
4. ненужная горутина в RemoveItemFromShoppingCart
Фактически это синхронная операция, искусственно усложнённая конкурентностью.
5. Глобальный канал removalChannel может приводить к логическим ошибкам
Пункты 4 и 5 ненужный overengineering
6. При удалении RemoveItemFromShoppingCart поиск идет по имени NameOfProduct
Если два товара имеют одинаковое имя, удалится только первый.
При удалении необходимо использовать уникальный id
7. Ошибки при расчете VAT
сумма товаров subTotal считается 2 раза
CalculateVAT - счтает только ндс, а потом прибавляет к повторно сложенному subtotal
8. Расчет суммы и вывод в 1 функции ListAndCalculateVAT
9. float32 — плохой выбор для денежных значений, т.к. имеет приблизительную точность
Для хранения денег используют:
- int64 - сумму в минимальных денежных единицах, (копейках или центах)
- специализированный тип денег или библиотеку, когда нужна более сложная арифметика
10. ставка НДС - глобальная переменная
Ставка может различаться по стране, категории товара и типу операции.
Лучше передавать налоговую ставку через конфигурацию или отдельную зависимость.
11. Заказ глобальная переменная, которую мы просто перезаписали
Никуда не сохраняется, никак не обрабатывается, никуда не пушится
Заказ и корзина ссылаются на один слайс Items, который может измениться после оформления заказа
12.

*/

/*
4-5
* sync.Mutex защищает общее изменяемое состояние от одновременного доступа.
* chan позволяет передавать данные и синхронизировать выполнение между горутинами.
* Каналы особенно полезны для передачи владения данными,
организации конвейеров и координации задач.
Не нужно создавать горутину и канал только ради того, чтобы вернуть результат операции.
Если работа синхронная и не требует конкурентного выполнения,
обычный вызов функции проще и безопаснее.
6
В реальной системе лучше различать сущность товара (ProductID), сам товар и позицию корзины.
Например, одна и та же модель товара может находиться в корзине в количестве трёх штук.

*/
