package xu

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type labeledContent struct {
	label   string
	content string
}

// AssertEqual checks if provided values are equal.
func AssertEqual(t *testing.T, actual interface{}, expected interface{}) error {
	if !reflect.DeepEqual(expected, actual) {
		msg := fmt.Sprintf("\nNot Equal:\n"+"  - Expected: %#v\n  - Received: %#v", expected, actual)
		return errors.New(msg)
	}
	return nil
}
