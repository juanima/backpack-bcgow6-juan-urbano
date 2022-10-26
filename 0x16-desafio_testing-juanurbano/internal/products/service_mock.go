package products

import (
	"fmt"
)

type MockService struct {
	DataMock []Product
        GetAllBySellerWasCalled bool
	Error    string
}

func (m *MockService) GetAllBySeller(sellerID string) ([]Product, error) {

        if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}

        m.GetAllBySellerWasCalled = true

        var updated bool
        var product Product
	for i := range m.DataMock {
		if m.DataMock[i].SellerID == sellerID {
			product = m.DataMock[i]
			updated = true
		}
	}

	if !updated {
		return []Product{}, fmt.Errorf("product sellerID sellerID %s not exists", sellerID)
	}

	return []Product{
                product,
        }, nil
}

