package ktest

import (
	"testing"

	"github.com/davelondon/ktest/assert"
)

func TestImports(t *testing.T) {
	if assert.Equal(t, 1, 1) != true {
		t.Error("Something is wrong.")
	}
}
