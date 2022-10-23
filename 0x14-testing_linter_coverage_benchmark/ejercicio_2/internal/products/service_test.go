package products

import (
        "fmt"
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ejercicio_2/pkg/store"
	"github.com/ejercicio_2/internal/domain"
)

func TestServiceIntegrationUpdate(t *testing.T) {
	// Arrange.
        publish := false
	database := []domain.Product{
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

	mockStorage := store.MockFileStoreIntegration{
		MockedData: database,
                ReadWasCalled: false,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act.
	product, err := service.Update("Caja de Galletas Bimbo", "red", "AC323", 1378, 32, 4.9, &publish) 

	// Assert.
	assert.Nil(t, err)
        assert.Equal(t, database[0], product)
}

func TestServiceIntegrationUpdateFail(t *testing.T) {

        // Arrange.
	expectedErr := errors.New("hello, i'm an error :D")

	mockStorage := store.MockFileStoreIntegration{
		MockedData: nil,
                ReadWasCalled: false,
		ErrOnWrite: nil,
		ErrOnRead:  errors.New("hello, i'm an error :D"),
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act.
        publish := false
        productToUpdate := domain.Product{
                Id:    1378,
                Name:  "Caja de galleticas oreo",
                Stock: 90,
                Color: "blue",
                Code:  "AC323",
                Price: 4.5,
                Publish: &publish,
        }
	
	product, err := service.Update(productToUpdate.Name, productToUpdate.Color, productToUpdate.Code, productToUpdate.Id, productToUpdate.Stock, productToUpdate.Price, productToUpdate.Publish) 

	// Assert.
        assert.Empty(t, product)
	assert.EqualError(t, err, expectedErr.Error())
}


func TestServiceIntegrationDelete(t *testing.T) {
	// Arrange.
	expectedErr := errors.New("Not Found Product")
        publish := false
	expectedDatabase := []domain.Product{
		{
			Id:    1378,
			Name:  "Caja de galleticas oreo",
                        Stock: 90,
			Color: "blue",
			Code:  "AC323",
			Price: 4.5,
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

	mockStorage := store.MockFileStoreIntegration{
		MockedData: initialDatabase,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act.
	errDelete := service.Delete(1399)
        product, errGetone := service.GetOne(1399)

	// Assert.
	assert.EqualError(t, errGetone, expectedErr.Error())
        assert.Nil(t, errDelete)
        assert.Empty(t, product)
        assert.Equal(t, len(initialDatabase) - 1, len(expectedDatabase))
}



func TestServiceIntegrationDeleteFail(t *testing.T) {
	// Arrange.
        idToDeleteProduct := 1399
        expectedErr := fmt.Errorf("product id %d not exists", idToDeleteProduct)

	mockStorage := store.MockFileStoreIntegration{
		MockedData:   nil,
		ErrOnRead:  nil,
		ErrOnWrite: errors.New("hello, i'm a little bug >:("),
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act.
	err := service.Delete(idToDeleteProduct)

	// Assert.
	assert.EqualError(t, err, expectedErr.Error())
}