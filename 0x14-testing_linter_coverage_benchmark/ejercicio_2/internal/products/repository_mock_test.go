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


func TestStoreGood(t *testing.T) {
	// Arrange.
        publish := false
	expected := []domain.Product{
		{
			Id:    1378,
			Name:  "Caja de galleticas oreo",
                        Stock: 90,
			Color: "blue",
			Code:  "AC323",
			Price: 4.5,
                        Publish: &publish,
		},
		{
			Id:    1399,
			Name:  "Desodorante Roxana",
                        Stock: 30,
			Color: "blue",
			Code:  "AC324",
			Price: 4.9,
                        Publish: &publish,
		},
		{
			Id:    1400,
			Name:  "Mexana",
                        Stock: 32,
			Color: "black",
			Code:  "AC329",
			Price: 4.9,
                        Publish: &publish,
                },
	}

	initialDatabase := []domain.Product{
		{
			Id:    1378,
			Name:  "Caja de galleticas oreo",
                        Stock: 90,
			Color: "blue",
			Code:  "AC323",
			Price: 4.5,
                        Publish: &publish,
		},
		{
			Id:    1399,
			Name:  "Desodorante Roxana",
                        Stock: 30,
			Color: "blue",
			Code:  "AC324",
			Price: 4.9,
                        Publish: &publish,
		},
	}

	storedb := &store.MockFileStore{
		MockedData: initialDatabase,
                ReadWasCalled: false,
	}

	repository := NewRepository(storedb)

	// Act.
	productToCreate := domain.Product{
		Id:    1400,
		Name:  "Mexana",
                Stock: 32,
		Color: "black",
		Code:  "AC329",
		Price: 4.9,
                Publish: &publish,
        }

        // Store(name, color, code string, stock int, price float64, publish *bool)
	result, err := repository.Store(productToCreate.Name, productToCreate.Color, productToCreate.Code, productToCreate.Stock, productToCreate.Price, productToCreate.Publish)

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, expected, storedb.MockedData)
	assert.Equal(t, productToCreate, result)
}

