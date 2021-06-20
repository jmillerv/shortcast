package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testPascalString = "HelloWorld"
	testCamelString = "helloTheWorldIsOnFire"
)

func TestAddSpacesToPascalString(t *testing.T) {
	s := AddSpacesToPascalString(testPascalString)
	assert.Equal(t, "Hello World", s)
}

func TestAddSpacesToCamelString(t *testing.T) {
	s := AddSpacesToCamelString(testCamelString)
	assert.Equal(t, "hello The World Is On Fire", s)
}
