package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func iRunTheAssertion(t *testing.T) {
	assert.NotZero(t, len(sampleData), "expected some len")
}