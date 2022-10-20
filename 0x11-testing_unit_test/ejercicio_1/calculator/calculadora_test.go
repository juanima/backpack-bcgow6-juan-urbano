package calculator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSubGood(t *testing.T) {
	// Arrange
	minuend := 10
	subtrahend := 5
	expected := 5

	// Act
	difference := Sub(minuend, subtrahend)

	// Assert
	assert.Equal(t, expected, difference, "El numero resultado: %d, es igual al expected: %d ", difference, expected)
}

func TestGoodSubNegativeMinuendSubtrahend(t *testing.T) {
	// Arrange
	minuend := -10
	subtrahend := -5
	expected := -5

	// Act
	difference := Sub(minuend, subtrahend)

	// Assert
	assert.Equal(t, expected, difference, "El numero resultado: %d, es igual al expected: %d ", difference, expected)
}

