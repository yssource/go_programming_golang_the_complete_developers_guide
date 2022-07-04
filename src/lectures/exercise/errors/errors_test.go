package timeparse

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{
			time: "19:00:42",
			ok:   true,
		},
		{
			time: "1:3:44",
			ok:   true,
		},
		{
			time: "bad",
			ok:   false,
		},
		{
			time: "1:-3:44",
			ok:   false,
		},
		{
			time: "00:59:59",
			ok:   true,
		},
		{
			time: "",
			ok:   false,
		},
		{
			time: "11:22",
			ok:   false,
		},
		{
			time: "aa:bb:cc",
			ok:   false,
		},
		{
			time: "5:23",
			ok:   false,
		},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}
}
