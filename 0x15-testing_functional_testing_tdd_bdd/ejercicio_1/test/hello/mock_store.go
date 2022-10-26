package hello

import (
        // "log"
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
                return errors.New("Ha pasado read!")
        }

	*products = mo.MockedData

	return nil
}

func (mo *MockFileStoreIntegration) Write(data interface{}) (err error) {

        if mo.ErrOnWrite != nil {
		return mo.ErrOnWrite
	}

        // Interfaces can be casted to a concrete type that implements it
	// products, ok := data.(*[]domain.Product)
        products, ok := data.([]domain.Product)
        if !ok {
                return errors.New("Ha pasado write!")
        }

        // log.Printf("------- %+v\n", products)

        // if len(products) == 0 {
        //         return errors.New("No hay products registrados!")
        // }

        // mo.MockedData = append(mo.MockedData, product)
	mo.MockedData = products

	return nil
}

