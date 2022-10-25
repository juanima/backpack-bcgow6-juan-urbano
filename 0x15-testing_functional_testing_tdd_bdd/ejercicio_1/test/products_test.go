package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	
	"github.com/ejercicio_2/test/hello"
        "github.com/ejercicio_2/cmd/server/handler"
	"github.com/ejercicio_2/internal/domain"
	"github.com/ejercicio_2/internal/products"
)

// Create an instance of gin Engine
func createServer(mockStore hello.MockFileStoreIntegration) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	repo := products.NewRepository(&mockStore)
	service := products.NewService(repo)
	control := handler.NewProduct(service)

	router := gin.Default()

	api := router.Group("/api/v1")
	pr := api.Group("/products")
        pr.PUT("/:id", control.Update())

	return router
}


// Create a http request "template"
func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestUpdateGood(t *testing.T) {
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

	mockStore := hello.MockFileStoreIntegration{
		MockedData: []domain.Product{
			beforeProduct,
		},
	}

	var resp domain.Product
        productToUpdate, err := json.Marshal(&afterUpdate)
        if err != nil {
            panic (err)
        }

	router := createServer(mockStore)
	req, rr := createRequestTest(http.MethodPut, "/api/v1/products/1379", string(productToUpdate))

        // Act
	router.ServeHTTP(rr, req)
	err = json.Unmarshal(rr.Body.Bytes(), &resp)

        // fmt.Printf("%+v\n", rr)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
	assert.Equal(t, expected, resp)
}

