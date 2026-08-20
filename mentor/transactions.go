/*
Нужно написать функцию которая восстановит итоговый баланс каждого пользователя.

Override == false, значения прибавляются
Override == true, значения перезаписываются
*/

package main

import (
	"fmt"
	"sort"
	"time"
)

type TransactionEvent struct {
	ID            string
	UserID        int
	DebitBalance  int
	CreditBalance int
	Timestamp     int64
	Override      bool
}

type UserBalance struct {
	Debit  int
	Credit int
}

func main() {
	transactions := map[string]TransactionEvent{
		"ab12cd34ef56": {
			ID:            "ab12cd34ef56",
			UserID:        1,
			DebitBalance:  500,
			CreditBalance: 200,
			Timestamp:     time.Now().Unix(),
			Override:      false,
		},
		"34ff78aa129b": {
			ID:            "34ff78aa129b",
			UserID:        1,
			DebitBalance:  750,
			CreditBalance: 300,
			Timestamp:     time.Now().Add(2 * time.Minute).Unix(),
			Override:      false,
		},
		"98aa21ff44de": {
			ID:            "98aa21ff44de",
			UserID:        1,
			DebitBalance:  1200,
			CreditBalance: 450,
			Timestamp:     time.Now().Add(4 * time.Minute).Unix(),
			Override:      true,
		},
		"aa99cc77bb66": {
			ID:            "aa99cc77bb66",
			UserID:        1,
			DebitBalance:  300,
			CreditBalance: 100,
			Timestamp:     time.Now().Add(6 * time.Minute).Unix(),
			Override:      true,
		},
		"1122ffee8899": {
			ID:            "1122ffee8899",
			UserID:        1,
			DebitBalance:  950,
			CreditBalance: 500,
			Timestamp:     time.Now().Add(8 * time.Minute).Unix(),
		},
		"dd1199aabb77": {
			ID:            "dd1199aabb77",
			UserID:        2,
			DebitBalance:  900,
			CreditBalance: 150,
			Timestamp:     time.Now().Add(10 * time.Minute).Unix(),
			Override:      false,
		},
		"66aa55cc4433": {
			ID:            "66aa55cc4433",
			UserID:        2,
			DebitBalance:  400,
			CreditBalance: 350,
			Timestamp:     time.Now().Add(12 * time.Minute).Unix(),
			Override:      true,
		},
		"ff3344dd5566": {
			ID:            "ff3344dd5566",
			UserID:        2,
			DebitBalance:  700,
			CreditBalance: 600,
			Timestamp:     time.Now().Add(14 * time.Minute).Unix(),
		},
		"bbccddeeff00": {
			ID:            "bbccddeeff00",
			UserID:        2,
			DebitBalance:  1100,
			CreditBalance: 800,
			Timestamp:     time.Now().Add(16 * time.Minute).Unix(),
			Override:      false,
		},
		"778899aabbcc": {
			ID:            "778899aabbcc",
			UserID:        2,
			DebitBalance:  500,
			CreditBalance: 200,
			Timestamp:     time.Now().Add(18 * time.Minute).Unix(),
		},
	}

	userBalance := RestoreBalance(transactions)

	fmt.Println(userBalance)
}

func RestoreBalance(log map[string]TransactionEvent) map[int]UserBalance {
	users := make(map[int]UserBalance)

	events := make([]TransactionEvent, 0, len(log))

	for _, value := range log {
		events = append(events, value)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Timestamp < events[j].Timestamp
	})

	for _, event := range events {
		balance := users[event.UserID]

		if event.Override {
			balance.Debit = event.DebitBalance
			balance.Credit = event.CreditBalance
		} else {
			balance.Debit += event.DebitBalance
			balance.Credit += event.CreditBalance
		}

		users[event.UserID] = balance
	}

	return users
}
