package file

import (
    "os"
    "encoding/json"
    "github.com/ejercicio_1/internal/service"
)

type File struct {
	path string
}

func (f *File) Read() (products []service.Product, err error) {
    
    data, err := os.ReadFile(f.path)
    if err != nil {
	return products, err
    }

    err = json.Unmarshal(data, &products)
    if err != nil { 
	return products, err
    }

    return products, nil
}

func (f *File) SetPath(path string) {
    f.path = path
}

