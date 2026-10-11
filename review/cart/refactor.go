package main

/*
import (
	"fmt"
	"net/smtp"
)

const vatRate = 10.0

type Product struct {
	NameOfProduct  string
	PriceOfProduct float32
}

type Order struct {
	Items []Product
}

type Cart struct {
	Items []Product
}

func NewCart() *Cart {
	return &Cart{
		Items: make([]Product, 0, 100),
	}
}

func NewOrder(cart Cart) *Order {
	return &Order{
		Items: cart.Items,
	}
}

func (c *Cart) AddItem(item Product) {
	c.Items = append(c.Items, item)
}

func (c *Cart) RemoveItem(nameOfItem string) bool {
	for i, item := range c.Items {
		if item.NameOfProduct == nameOfItem {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			return true
		}
	}
	return false
}

func (c *Cart) CalculateVAT() float64 {
	subTotal := 0.0
	for _, item := range c.Items {
		price := float64(item.PriceOfProduct)
		subTotal += price
	}

	vatAmount := subTotal * (vatRate / 100)

	return vatAmount
}

func (c *Cart) ListAndCalculateVAT() float64 {
	subTotal := 0.0
	for _, item := range c.Items {
		fmt.Printf("Item Name: %s, Item Price: $%.2f\n", item.NameOfProduct, item.PriceOfProduct)
		subTotal += float64(item.PriceOfProduct)
	}

	vatAmount := c.CalculateVAT()
	totalWithVAT := subTotal + vatAmount

	fmt.Printf("The total price with VAT: $%.2f\n", totalWithVAT)

	return totalWithVAT
}

func CreateOrder(cart Cart, customerEmail string) *Order {
	totalPrice := cart.ListAndCalculateVAT()

	orderDetails := "Thank you for your order!\n\n"
	orderDetails += "Here are the items in your order:\n"

	for _, item := range cart.Items {
		orderDetails += fmt.Sprintf("- %s: $%.2f\n", item.NameOfProduct, item.PriceOfProduct)
	}

	orderDetails += fmt.Sprintf("\nTotal (including VAT): $%.2f\n", totalPrice)

	subject := "Your Order Confirmation"

	err := SendOrderConfirmationEmail(customerEmail, subject, orderDetails)

	order := NewOrder(cart)

	if err != nil {
		fmt.Println("Failed to send order confirmation email:", err)
	}

	return order
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

*/
