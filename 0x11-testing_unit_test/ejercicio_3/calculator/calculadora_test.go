package calculator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDivideGood(t *testing.T) {
	// Arrange
        quotient, divisor:= 10, 5
	expected := 2

	// Act
	dividend, err := Divide(quotient, divisor)

	// Assert
	assert.Equal(t, expected, dividend)
	assert.Nil(t, err)
}

func TestDivideFail(t *testing.T) {
	// Arrange
        quotient, divisor:= 10, 0
	expected := 0

	// Act
	dividend, err := Divide(quotient, divisor)

	// Assert
	assert.Equal(t, expected, dividend, "El divisor no puede ser: %d", divisor)
	assert.NotNil(t, err)
}
