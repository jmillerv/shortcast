package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testString = "HelloWorld"
)

func TestAddSpacesToSentence(t *testing.T) {
	s := AddSpacesToPascalString(testString)
	assert.Equal(t, "Hello World", s)
}
