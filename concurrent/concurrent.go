package concurrent

// Execute will process all inputs concurrently by calling the function passed in the arguments.
// The number of goroutines that are used in the concurrent execution could be specified in the numOfRoutines parameter.
// The execution follows fan-out and then fan-in pattern, in which multiple processes are run concurrently, then each
// outputs are gathered and appended to a single slice at the end of the execution. The slice is not guaranteed to have
// the one-to-one order as the input, so it is advised to not rely on the output slice order.
func Execute[TypeIn any, TypeOut any](numOfRoutines int, inputs []TypeIn, process func(input TypeIn) TypeOut) []TypeOut {
	inputChannel := make(chan TypeIn)
	outputChannel := make(chan TypeOut)

	// spawn workers
	for i := 0; i < numOfRoutines; i++ {
		go func(inputChan chan TypeIn, outputChan chan TypeOut) {
			for input := range inputChan {
				outputChan <- process(input)
			}
		}(inputChannel, outputChannel)
	}

	// distribute inputs
	go func(inputs []TypeIn, inputChannels chan TypeIn) {
		for _, input := range inputs {
			inputChannel <- input
		}
		close(inputChannel)
	}(inputs, inputChannel)

	// wait for outputs
	outputs := []TypeOut{}
	doneCount := 0
	for {
		outputs = append(outputs, <-outputChannel)
		doneCount++
		if doneCount >= len(inputs) {
			close(outputChannel)
			break
		}
	}
	return outputs
}
