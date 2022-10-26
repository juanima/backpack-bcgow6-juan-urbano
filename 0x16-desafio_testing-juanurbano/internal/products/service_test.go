package products

import (
	// "errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit test Service func GetAlBySeller Success
func TestServiceGetAllBySellerGood(t *testing.T) {
	// arrange
        sellerID := "FEX112AC"
        database := []Product{
                {
        		ID:          "mock",
        		SellerID:    sellerID,
        		Description: "generic product",
        		Price:       123.55,
                },
                {
                        ID:          "mock_2",
                        SellerID:    "FIX113AE",
                        Description: "blue",
                        Price:       4.7,
                },
        }

	mockRepository := MockRepository{
		DataMock: database,
	}

	// act
	service := NewService(&mockRepository)
	result, err := service.GetAllBySeller(sellerID)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, sellerID, result[0].SellerID)
}


// Unit test Service func GetAlBySeller id not found
func TestServiceGetAllBySellerFail(t *testing.T) {
	// arrange
        sellerID := "FEX112AC"
        errorMessage := fmt.Sprintf("No encontró sellerID = %s", sellerID)

	mockRepository := MockRepository{
		DataMock: []Product{},
                Error: errorMessage,
	}

	// act
	service := NewService(&mockRepository)
	result, err := service.GetAllBySeller(sellerID)

	// assert
	assert.NotNil(t, err)
        assert.Nil(t, result)
}

