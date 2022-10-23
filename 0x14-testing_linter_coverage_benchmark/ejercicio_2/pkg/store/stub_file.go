package store

import (
        "errors"
        "github.com/ejercicio_2/internal/domain"
)

type StubFileStore struct{
        MockedData []domain.Product
}


func (st StubFileStore) Read(data interface{}) error {

        a, ok := data.(*[]domain.Product)
        if !ok {
                return errors.New("Ha pasado!")
        }

        *a = st.MockedData

        return nil
}

func (st StubFileStore) Write(data interface{}) error {
        return nil
}


