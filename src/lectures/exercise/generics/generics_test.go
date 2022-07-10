package generics

import (
	"fmt"
	"testing"
)

type my_test[T Number] struct {
	value, min, max T
	expected        T
}

func TestClamp(t *testing.T) {
	var tableInt32 []my_test[int32] = []my_test[int32]{
		{value: 1, min: 0, max: 3, expected: 1},
		{value: 0, min: 1, max: 3, expected: 1},
		{value: 4, min: 1, max: 3, expected: 3},
	}

	for _, s := range tableInt32 {
		v := Clamp(s.value, s.min, s.max)
		if v != s.expected {
			t.Errorf(fmt.Sprintf("invalid result:\n\tExpected: %v\n\tGot: %v", s.expected, v))
		}
	}

	var tableFloat64 []my_test[float64] = []my_test[float64]{
		{value: 1.2, min: 0.0, max: 3.0, expected: 1.2},
		{value: 0.5, min: 1.0, max: 3.2, expected: 1},
		{value: 4.2, min: 1.0, max: 3.3, expected: 3.3},
	}

	for _, s := range tableFloat64 {
		v := Clamp(s.value, s.min, s.max)
		if v != s.expected {
			t.Errorf(fmt.Sprintf("invalid result:\n\tExpected: %v\n\tGot: %v", s.expected, v))
		}
	}

	var tableVeolcity []my_test[Velocity] = []my_test[Velocity]{
		{value: 1.2, min: 0.0, max: 3.0, expected: 1.2},
		{value: 0.5, min: 1.0, max: 3.2, expected: 1},
		{value: 4.2, min: 1.0, max: 3.3, expected: 3.3},
	}

	for _, s := range tableVeolcity {
		v := Clamp(s.value, s.min, s.max)
		if v != s.expected {
			t.Errorf(fmt.Sprintf("invalid result:\n\tExpected: %v\n\tGot: %v", s.expected, v))
		}
	}

	var tableDistance []my_test[Distance] = []my_test[Distance]{
		{value: 1, min: 0, max: 3, expected: 1},
		{value: 0, min: 1, max: 3, expected: 1},
		{value: 4, min: 1, max: 3, expected: 3},
	}

	for _, s := range tableDistance {
		v := Clamp(s.value, s.min, s.max)
		if v != s.expected {
			t.Errorf(fmt.Sprintf("invalid result:\n\tExpected: %v\n\tGot: %v", s.expected, v))
		}
	}
}
