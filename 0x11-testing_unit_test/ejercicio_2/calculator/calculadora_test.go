package calculator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSortSliceGood(t *testing.T) {
	// Arrange
        arr := []int{30, 13, 8, 9}
        expected := []int{8, 9, 13, 30}

	// Act
	sortArr := SortSlice(arr)

	// Assert
	assert.Equal(t, expected, sortArr)
}

func TestSortSliceGoodSameNumber(t *testing.T) {
	// Arrange
        arr := []int{30, 30, 30, 30}
        expected := []int{30, 30, 30, 30}

	// Act
	sortArr := SortSlice(arr)

	// Assert
	assert.Equal(t, expected, sortArr)
}

