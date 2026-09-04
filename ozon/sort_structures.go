package main

import (
	"fmt"
	"sort"
)

var priorities = map[string]int{
	"Epic":     1,
	"Story":    2,
	"Task":     3,
	"Sub-Task": 4,
}

type Issue struct {
	ID   int
	Type string
}

func main() {
	issues := []Issue{
		{ID: 2, Type: "Sub-Task"},
		{ID: 1, Type: "Sub-Task"},
		{ID: 3, Type: "Task"},
		{ID: 4, Type: "Story"},
		{ID: 5, Type: "Epic"},
	}

	sort.Slice(issues, func(i, j int) bool {
		priorityA := priorities[issues[i].Type]
		priorityB := priorities[issues[j].Type]

		if priorityA != priorityB {
			return priorityA < priorityB
		}

		return issues[i].ID < issues[j].ID
	})

	fmt.Println(issues)
}
