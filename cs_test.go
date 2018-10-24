package cs_test

import (
	"testing"

	"github.com/Konstantin8105/cs"
)

func TestAll(t *testing.T) {
	tcs := []struct {
		name string
		f    func(*testing.T)
	}{
		{"Todo", cs.Todo},
		{"Debug", cs.Debug},
		//	{"Os", cs.Os}, // ignore
	}
	for _, tc := range tcs {
		t.Run(tc.name, tc.f)
	}
}
