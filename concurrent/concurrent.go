package concurrent

import "sync"

// Execute will process all inputs concurrently by calling the function passed in the arguments.
// The number of goroutines that are used in the concurrent execution could be specified in the numOfRoutines parameter.
// The execution follows fan-out and then fan-in pattern, in which multiple processes are run concurrently, then each
// outputs are gathered and appended to a single slice at the end of the execution. The slice is not guaranteed to have
// the one-to-one order as the input, so it is advised to not rely on the output slice order.
func Execute[TypeIn any, TypeOut any](numOfRoutines int, inputs []TypeIn, process func(input TypeIn) TypeOut) []TypeOut {
	inputChannel := make(chan TypeIn)
	outputChannel := make(chan TypeOut)
	var wg sync.WaitGroup

	// spawn workers
	for i := 0; i < numOfRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for input := range inputChannel {
				outputChannel <- process(input)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(outputChannel)
	}()

	// distribute inputs
	go func() {
		for _, input := range inputs {
			inputChannel <- input
		}
		close(inputChannel)
	}()

	// wait for outputs
	outputs := []TypeOut{}
	for output := range outputChannel {
		outputs = append(outputs, output)
	}
	return outputs
}
