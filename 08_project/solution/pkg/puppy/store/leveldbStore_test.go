package store

import (
	"testing"

	puppy "github.com/anz-bank/go-course/08_project/solution/pkg/puppy"
)

func TestDbErrorPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Panic expected")
		}
	}()
	err := puppy.Errorf(puppy.ErrInvalidInputCode, "test invalid input")
	dbErrorPanic(err)
}
