package countdown

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountdown(t *testing.T) {
	testCases := []TestCase{
		{
			Name:           "testing countdown from 5",
			Input:          5,
			ExpectedOutput: "5, 4, 3, 2, 1 Done!",
		},
		{
			Name:           "testing countdown from 4",
			Input:          4,
			ExpectedOutput: "4, 3, 2, 1 Done!",
		},
		{
			Name:           "testing countdown from 3",
			Input:          3,
			ExpectedOutput: "3, 2, 1 Done!",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			Countdown(buffer, tc.Input)
			result := buffer.String()
			if result != tc.ExpectedOutput {
				t.Errorf(fmt.Sprintf("expected %s but got %s", tc.ExpectedOutput, result))
			}
		})
	}

}

type TestCase struct {
	Name           string
	Input          int
	ExpectedOutput string
}
