package products

import (
	"testing"
	"github.com/ejercicio_2/internal/domain"
	"github.com/ejercicio_2/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	expected := []domain.Product{
		{
			Id:    1378,
			Name:  "beer",
			Color: "blue",
			Code:  "AC323",
			Price: 4.5,
		},
	}

	// Arrange
	// db := store.New(store.FileType, "./products.json")
	storedb := &store.StubFileStore{
		MockedData: expected,
	}

	repo := NewRepository(storedb)

	// Act
	products, err := repo.GetAll()

	// Assert
	assert.Equal(t, expected, products)
	assert.Nil(t, nil, err)
}
