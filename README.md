# gorutin
Go concurrency generics

## Documentation
https://pkg.go.dev/github.com/raymondhartoyo/gorutin

## Usage

### Processing a set of inputs concurrently

Given a set of inputs and a function to be run concurrently, you can call `concurrent.Execute` as shown in the following example:

```
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

outputs := concurrent.Execute(4, inputs, process)
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
```

## Release Information

### v0.1
- This is still an unstable release, public APIs (function signatures might be updated at the future release).
- See [CHANGELOG.md](CHANGELOG.md)

## Future milestones
- Context cancellation
- Native timeout support

## Contributing
If you find any bug or suggestion, please raise it in this project's issues page.
