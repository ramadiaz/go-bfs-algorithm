package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreadthFirstSearch(t *testing.T) {
	tree := map[string]Node{
		"John":     {Value: "John", Neighbors: []string{"George", "Sam", "Edward"}},
		"George":   {Value: "George", Neighbors: []string{"Richard", "John"}},
		"Sam":      {Value: "Sam", Neighbors: []string{"Richard", "Briana"}},
		"Edward":   {Value: "Edward", Neighbors: []string{"Anett", "Shaun"}},
		"Richard":  {Value: "Richard", Neighbors: []string{"Franklin"}},
		"Briana":   {Value: "Briana", Neighbors: []string{"Lynsey", "Karen"}},
		"Anett":    {Value: "Anett", Neighbors: []string{"Wilson"}},
		"Shaun":    {Value: "Shaun", Neighbors: []string{}},
		"Franklin": {Value: "Franklin", Neighbors: []string{}},
		"Lynsey":   {Value: "Lynsey", Neighbors: []string{}},
		"Karen":    {Value: "Karen", Neighbors: []string{}},
		"Wilson":   {Value: "Wilson", Neighbors: []string{}},
	}

	tests := []struct {
		name     string
		input    map[string]Node
		start    string
		end      string
		expected string
	}{
		{
			name:     "john_anett_example",
			input:    tree,
			start:    "John",
			end:      "Anett",
			expected: "John->Edward->Anett",
		},
		{
			name:     "george_john_example",
			input:    tree,
			start:    "George",
			end:      "John",
			expected: "George->John",
		},
		{
			name:     "roman_root_not_found",
			input:    tree,
			start:    "Roman",
			end:      "Anett",
			expected: "not_found",
		},
		{
			name:     "john_lynsey_example",
			input:    tree,
			start:    "John",
			end:      "Lynsey",
			expected: "John->Sam->Briana->Lynsey",
		},
		{
			name:     "emili_target_not_found",
			input:    tree,
			start:    "John",
			end:      "Emili",
			expected: "not_found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := BreadthFirstSearch(test.input, test.start, test.end)
			log.Println("RESULT FOR ", test.name, " IS ", result)

			if !assert.Equal(t, test.expected, result) {
				t.FailNow()
			}
		})
	}
}
