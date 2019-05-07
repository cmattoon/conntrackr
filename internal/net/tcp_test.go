package net

import (
	"testing"
)

func TestIsTCPState(t *testing.T) {
	for _, tc := range []struct {
		state  string
		result bool
	}{
		{
			state:  "ESTABLISHED",
			result: true,
		},
		{
			state:  "TIME_WAIT",
			result: true,
		},
		{
			state:  "Puerto Rico",
			result: false,
		},
		{
			state:  "Washington, DC",
			result: false,
		},
	} {
		if tc.result != IsTCPState(tc.state) {
			t.Fail("wrong")
		}
	}
}
