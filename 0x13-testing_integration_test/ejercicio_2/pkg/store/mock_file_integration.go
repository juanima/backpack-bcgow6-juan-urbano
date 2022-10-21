package store

import (
        "errors"
        "github.com/ejercicio_2/internal/domain"
)


type MockFileStoreIntegration struct {
        MockedData []domain.Product
        ReadWasCalled bool
	ErrOnWrite error
	ErrOnRead  error
}

func (mo *MockFileStoreIntegration) Read(data interface{}) (err error) {

	if mo.ErrOnRead != nil {
		return mo.ErrOnRead
	}

        mo.ReadWasCalled = true

	products, ok := data.(*[]domain.Product)
        if !ok {
                return errors.New("Ha pasado!")
        }

	*products = mo.MockedData

	return nil
}

func (mo *MockFileStoreIntegration) Write(data interface{}) (err error) {
	
        if mo.ErrOnWrite != nil {
		return mo.ErrOnWrite
	}

	product, ok := data.(domain.Product)
        if !ok {
                return errors.New("Ha pasado!")
        }
	
        mo.MockedData = append(mo.MockedData, product)

	return nil
}

