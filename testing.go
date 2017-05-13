package xu

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Errorf(format string, args ...interface{})
}

type labeledContent struct {
	label   string
	content string
}

// AssertEqual ...
func AssertEqual(t *testing.T, actual interface{}, expected interface{}) error {
	if !reflect.DeepEqual(expected, actual) {
		msg := fmt.Sprintf("\nNot Equal:\n"+"  - Expected: %#v\n  - Received: %#v", expected, actual)
		return errors.New(msg)
	}
	return nil
}
