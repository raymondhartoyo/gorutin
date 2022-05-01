package concurrent

// Execute will process all inputs concurrently by calling the function passed in the arguments.
// The number of goroutines that are used in the concurrent execution could be specified in the numOfRoutines parameter.
// The execution follows fan-out and then fan-in pattern, in which multiple processes are run concurrently, then each
// outputs are gathered and appended to a single slice at the end of the execution. The slice is not guaranteed to have
// the one-to-one order as the input, so it is advised to not rely on the output slice order.
func Execute[TypeIn any, TypeOut any](numOfRoutines int, inputs []TypeIn, process func(input TypeIn) TypeOut) []TypeOut {
	inputChannels := make([](chan TypeIn), numOfRoutines)
	outputChannels := make([](chan TypeOut), numOfRoutines)
	for i := 0; i < numOfRoutines; i++ {
		inputChannels[i] = make(chan TypeIn)
		outputChannels[i] = make(chan TypeOut)
	}

	// spawn workers
	for i := 0; i < numOfRoutines; i++ {
		inputChannel := inputChannels[i]
		outputChannel := outputChannels[i]
		go func(inputChan chan TypeIn, outputChan chan TypeOut) {
			defer close(outputChan)
			for input := range inputChan {
				outputChan <- process(input)
			}
		}(inputChannel, outputChannel)
	}

	// distribute inputs
	go func(inputs []TypeIn, inputChannels [](chan TypeIn)) {
		for i, input := range inputs {
			channel := i % numOfRoutines
			inputChannels[channel] <- input
		}
		for _, inputChan := range inputChannels {
			close(inputChan)
		}
	}(inputs, inputChannels)

	// wait for outputs
	outputs := []TypeOut{}
	for {
		closedCount := 0
		for i := 0; i < numOfRoutines; i++ {
			select {
			case o, open := <-outputChannels[i]:
				if !open {
					closedCount++
					continue
				}
				outputs = append(outputs, o)
			default:
				continue
			}
		}
		if closedCount >= numOfRoutines {
			break
		}
	}
	return outputs
}
