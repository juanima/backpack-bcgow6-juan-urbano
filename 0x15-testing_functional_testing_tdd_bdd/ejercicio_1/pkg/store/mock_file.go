package store

import (
        "errors"
        "github.com/ejercicio_2/internal/domain"
)

type MockFileStore struct{
        MockedData []domain.Product
        ReadWasCalled bool
}

func (mo *MockFileStore) Read(data interface{}) error {

        mo.ReadWasCalled = true
        products, ok := data.(*[]domain.Product)
        if !ok {
                return errors.New("Ha pasado!")
        }

        *products = mo.MockedData

        return nil
}

func (mo *MockFileStore) Write(data interface{}) error {

	// products, ok := data.(*[]domain.Product)
	products, ok := data.([]domain.Product)
        if !ok {
                return errors.New("Ha pasado write!")
        }

        // mo.MockedData = append(mo.MockedData, product)
	 mo.MockedData = products

        return nil
}


