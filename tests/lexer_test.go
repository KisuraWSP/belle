package tests

import (
	"testing"

	l "github.com/KisuraWSP/belle/src/lexer"
	"github.com/stretchr/testify/assert"
)

func TestAssert(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(l.Lexer(" "), " ", "Invalid")
}
