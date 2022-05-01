package concurrent_test

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
	"time"

	. "github.com/raymondhartoyo/gorutin/concurrent"
	"github.com/stretchr/testify/assert"
)

var (
	valueForErrorCase string = "value for error case"
)

type testInput struct {
	value string
}

type testOutput struct {
	value string
	err   error
}

func sleepForARandomTime() {
	random, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}
	randomSleepMs := random.Int64()
	time.Sleep(time.Duration(randomSleepMs) * time.Millisecond)
}

func testProcess(in testInput) testOutput {
	sleepForARandomTime()
	if in.value == valueForErrorCase {
		return testOutput{err: fmt.Errorf("error when processing %s", valueForErrorCase)}
	}
	return testOutput{value: fmt.Sprintf("correctly processing input for %s", in.value)}
}

func TestExecute(t *testing.T) {
	d1 := testInput{value: "value1"}
	d2 := testInput{value: "value2"}
	d3 := testInput{value: "value3"}
	d4 := testInput{value: "value4"}
	d5 := testInput{value: "value5"}
	d6 := testInput{value: "value6"}
	d7 := testInput{value: "value7"}
	d8 := testInput{value: valueForErrorCase}
	d9 := testInput{value: "value8"}

	oddInputs := []testInput{
		d1, d2, d3, d4, d5, d6, d7, d8, d9,
	}
	expectedOddOutputs := []testOutput{}
	for _, i := range oddInputs {
		expectedOddOutputs = append(expectedOddOutputs, testProcess(i))
	}

	evenInputs := []testInput{
		d2, d3, d4, d5, d6, d7, d8, d9,
	}
	expectedEvenOutputs := []testOutput{}
	for _, i := range evenInputs {
		expectedEvenOutputs = append(expectedEvenOutputs, testProcess(i))
	}

	testCases := []struct {
		name            string
		numOfRoutines   int
		inputs          []testInput
		expectedOutputs []testOutput
	}{
		{
			name:            "odd number of inputs with number of routines equal number of inputs",
			numOfRoutines:   9,
			inputs:          oddInputs,
			expectedOutputs: expectedOddOutputs,
		},
		{
			name:            "odd number of inputs with number of routines less than number of inputs",
			numOfRoutines:   7,
			inputs:          oddInputs,
			expectedOutputs: expectedOddOutputs,
		},
		{
			name:            "odd number of inputs with number of routines more than number of inputs",
			numOfRoutines:   14,
			inputs:          oddInputs,
			expectedOutputs: expectedOddOutputs,
		},
		{
			name:            "even number of inputs with number of routines equal number of inputs",
			numOfRoutines:   8,
			inputs:          evenInputs,
			expectedOutputs: expectedEvenOutputs,
		},
		{
			name:            "even number of inputs with number of routines less than number of inputs",
			numOfRoutines:   1,
			inputs:          evenInputs,
			expectedOutputs: expectedEvenOutputs,
		},
		{
			name:            "even number of inputs with number of routines more than number of inputs",
			numOfRoutines:   16,
			inputs:          evenInputs,
			expectedOutputs: expectedEvenOutputs,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			outputs := Execute(tc.numOfRoutines, tc.inputs, testProcess)
			assert.ElementsMatch(t, tc.expectedOutputs, outputs)
		})
	}
}
