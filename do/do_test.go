package do

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Do(t *testing.T) {
    tests := map[string] struct {
		inputString string
		inputInt int
		inputBool bool
		expectedResult string
		expectedError string
	}{
		"success,b true": {
			inputString: "a",
			inputInt: 2,
			inputBool: true,
			expectedResult: "[2A]",
			expectedError: "",
		},
		"success,b false": {
			inputString: "b",
			inputInt: 8,
			inputBool: false,
			expectedResult: "[8b]",
			expectedError: "",
		},
		"success,len(s) < 2": {
			inputString: "c",
			inputInt: 34,
			inputBool: true,
			expectedResult: "C",
			expectedError: "",
		},
		"invalid parameter s string": {
			inputString: "x",
			inputInt: 2,
			inputBool: true,
			expectedResult: "",
			expectedError: "invalid s",
		},
		"invalid parametr i int": {
			inputString: "a",
			inputInt: 7,
			inputBool: true,
			expectedResult: "",
			expectedError: "invalid s",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(tt *testing.T) {
			result, err := Do(tc.inputString, tc.inputInt, tc.inputBool)
			assert.Equal(t, tc.expectedResult, result)

			if tc.expectedError != "" {
				assert.EqualError(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}