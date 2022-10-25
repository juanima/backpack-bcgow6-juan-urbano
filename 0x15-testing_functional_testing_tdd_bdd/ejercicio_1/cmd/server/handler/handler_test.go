package handler

import (
        // "fmt"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/ejercicio_2/internal/domain"
	"github.com/ejercicio_2/test/hello"
)

func createServer(mockService hello.MockService) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	handler := NewProduct(&mockService)
	router := gin.Default()

	pr := router.Group("/products")
	pr.POST("", handler.Store())
	pr.GET("", handler.GetAll())
	pr.PUT("/:id", handler.Update())

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}


func TestGetAllStatusOk(t *testing.T) {
	// Arrange
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
                        Date: "",
		},
	}

	mockService:= hello.MockService{
		DataMock: database,
	}

        var resp []domain.Product

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/products", "")

	// Act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)

	// Assert
	assert.Nil(t, err)
        assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, mockService.DataMock, resp)
}


func TestGetAllStatusNotFoundError(t *testing.T) {
	// Arrange
	mockService := hello.MockService{
		DataMock: []domain.Product{},
		Error:    "",
	}

	var resp map[string]string
	expected := map[string]string{
                "code": "404",
		"error": "no hay productos registrados",
	}

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/products", "")

	// Act
	router.ServeHTTP(rr, req)

	// Assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
        assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, expected, resp)
}

func TestStore(t *testing.T) {
	// Arrange
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
                        Date: "",
		},
	}

	expected := domain.Product{
                Id: 1379,
		Name:  "Beer",
                Stock: 90,
		Color: "blue",
		Code:  "AC353",
		Price: 4.5,
                Publish: &publish,
                Date: "",
	}

	requestBodyJSON := domain.Product{
                Id: 1379,
		Name:  "Beer",
                Stock: 90,
		Color: "blue",
		Code:  "AC353",
		Price: 4.5,
                Publish: &publish,
                Date: "December",
	}

	mockService:= hello.MockService{
		DataMock: database,
	}

	var resp domain.Product

        // Act
	router := createServer(mockService)

        productToSave, err := json.Marshal(&requestBodyJSON)
        if err != nil {
            panic (err)
        }

	req, rr := createRequestTest(http.MethodPost, "/products", string(productToSave))
	router.ServeHTTP(rr, req)
	err = json.Unmarshal(rr.Body.Bytes(), &resp)

        // fmt.Printf("%+v\n", rr)

	// Assert
	assert.Nil(t, err)
        assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, expected, resp)
}


func TestPutOk(t *testing.T) {
	// Arrange
        publish := false
	beforeProduct := domain.Product{
                Id: 1379,
		Name:  "Beer",
                Stock: 90,
		Color: "blue",
		Code:  "AC353",
		Price: 4.5,
                Publish: &publish,
                Date: "January",
	}

        afterUpdate := domain.Product{
                Id: 1379,
		Name:  "Beer",
                Stock: 99,
		Color: "blue",
		Code:  "AC353",
		Price: 4.7,
                Publish: &publish,
                Date: "January",
        }

	expected := domain.Product{
                Id: 1379,
		Name:  "Beer",
                Stock: 99,
		Color: "blue",
		Code:  "AC353",
		Price: 4.7,
                Publish: &publish,
                Date: "",
	}

	mockService := hello.MockService{
		DataMock: []domain.Product{
			beforeProduct,
		},
		Error:    "",
	}

	var resp domain.Product
        productToUpdate, err := json.Marshal(&afterUpdate)
        if err != nil {
            panic (err)
        }

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodPut, "/products/1379", string(productToUpdate))

        // Act
	router.ServeHTTP(rr, req)
	err = json.Unmarshal(rr.Body.Bytes(), &resp)

        // fmt.Printf("%+v\n", rr)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
	assert.Equal(t, expected, resp)
}


