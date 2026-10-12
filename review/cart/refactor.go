package main

import (
	"sync"
)

type ProductR struct {
	Name  string
	Price int64
}

type OrderR struct {
	Items []ProductR
}

type Cart struct {
	mu      sync.Mutex
	items   []ProductR
	vatRate float64
}

func NewCart(vatRate float64) *Cart {
	return &Cart{
		items:   make([]ProductR, 0),
		vatRate: vatRate,
	}
}

func (c *Cart) AddItem(item ProductR) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = append(c.items, item)
}

func (c *Cart) RemoveItem(name string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i, item := range c.items {
		if item.Name == name {
			c.items = append(c.items[:i], c.items[i+1:]...)
			return true
		}
	}

	return false
}

func (c *Cart) CalculateTotals() (
	subtotal int64,
	vat int64,
	total int64,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, item := range c.items {
		subtotal += item.Price
	}

	vat = int64(float64(subtotal)*c.vatRate/100 + 0.5)
	total = subtotal + vat

	return subtotal, vat, total
}

func (c *Cart) CreateOrder() OrderR {
	c.mu.Lock()
	defer c.mu.Unlock()

	orderItems := append([]ProductR(nil), c.items...)

	return OrderR{
		Items: orderItems,
	}
}
