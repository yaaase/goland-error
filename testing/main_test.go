package testing

import (
	"testing"
)

var sampleData = "I am a string defined in the _test.go file"

func TestThatCallsAnotherFunctionWhichReliesOnItsVariables(t *testing.T) {
	iRunTheAssertion(t)
}