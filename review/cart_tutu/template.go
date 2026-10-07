// В этом коде описана логика работы корзины,
// в которую можно добавлять продукты, убирать их, а также создавать заказ.

// Нужно провести Code Review и предложить улучшения по данному функционалу.

package cart

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
