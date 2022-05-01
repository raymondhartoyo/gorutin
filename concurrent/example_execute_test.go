package concurrent_test

import (
	"fmt"
	"sort"

	. "github.com/raymondhartoyo/gorutin/concurrent"
)

func ExampleExecute() {
	inputs := []string{
		"Alice",
		"Bob",
		"John",
		"Doe",
		"Sarah",
		"Susan",
	}

	process := func(input string) string {
		return fmt.Sprintf("Hi %s", input)
	}

	outputs := Execute(4, inputs, process)
	sort.Strings(outputs)
	for _, output := range outputs {
		fmt.Println(output)
	}

	// Output:
	// Hi Alice
	// Hi Bob
	// Hi Doe
	// Hi John
	// Hi Sarah
	// Hi Susan
}
