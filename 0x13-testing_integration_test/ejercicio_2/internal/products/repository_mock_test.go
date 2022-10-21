package products

import (
	"testing"
	"github.com/ejercicio_2/internal/domain"
	"github.com/ejercicio_2/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestUpdateNameGood(t *testing.T) {

        expectedNameUpdate := "After Update"
	productsExpected := []domain.Product{
		{
			Id:    1378,
			Name:  "Before Update",
			Color: "blue",
			Code:  "AC323",
			Price: 4.5,
		},
	}

	// Arrange
	// db := store.New(store.FileType, "./products.json")
	storedb := &store.MockFileStore{
		MockedData: productsExpected,
                ReadWasCalled: false,
	}

	repo := NewRepository(storedb)

	// Act
	product, err := repo.UpdateName(1378, expectedNameUpdate)

	// Assert
        assert.True(t, storedb.ReadWasCalled)
	assert.Equal(t, product.Name, expectedNameUpdate)
	assert.Nil(t, nil, err)
}


func TestUpdateNameBad(t *testing.T) {

        expectedNameUpdate := ""
	productExpected := domain.Product{}

	// Arrange
	// db := store.New(store.FileType, "./products.json")
	storedb := &store.MockFileStore{
		// MockedData: productsExpected,
                ReadWasCalled: false,
	}

	repo := NewRepository(storedb)

	// Act
	product, err := repo.UpdateName(1378, expectedNameUpdate)

	// Assert
        assert.True(t, storedb.ReadWasCalled)
	assert.Equal(t, product, productExpected)
	assert.Error(t, err)
}
