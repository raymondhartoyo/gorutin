package concurrent_test

import (
	"errors"
	"fmt"
	"sort"

	. "github.com/raymondhartoyo/gorutin/concurrent"
)

var (
	names = []string{
		"Alice",
		"Bob",
		"John",
		"Doe",
		"Sarah",
		"Susan",
	}
)

type greetingCard struct {
	greeting string
	err      error
}

func (g greetingCard) str() string {
	if g.err != nil {
		return fmt.Sprintf("err: %s", g.err)
	}
	return g.greeting
}

func createGreetingCard(name string) greetingCard {
	if name == "John" {
		return greetingCard{
			err: errors.New("no greeting for John"),
		}
	}
	return greetingCard{
		greeting: fmt.Sprintf("hi %s", name),
	}
}

// Example (OutputWithError) demonstrates a process that could return an error.
func ExampleExecute_outputWithError() {
	greetingCards := Execute(4, names, createGreetingCard)

	greetings := []string{}
	for _, card := range greetingCards {
		greetings = append(greetings, card.str())
	}
	sort.Strings(greetings)
	for _, greeting := range greetings {
		fmt.Println(greeting)
	}

	// Output:
	// err: no greeting for John
	// hi Alice
	// hi Bob
	// hi Doe
	// hi Sarah
	// hi Susan
}
