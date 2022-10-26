package products

import (
        // "fmt"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockService MockService) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	handler := NewHandler(&mockService)
	router := gin.Default()

	pr := router.Group("/api/v1")
	pr.GET("/products", handler.GetProducts)

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

// 
func TestHandlerGetProductsGood(t *testing.T) {
	// arrange
        database := []Product{
                {
        		ID:          "mock",
        		SellerID:    "FEX112AC",
        		Description: "generic product",
        		Price:       123.55,
                },
                {
                        ID:          "mock_2",
                        SellerID:    "FIX113AE",
                        Description: "blue",
                        Price:       4.7,
                },
        }

        mockService:= MockService{
		DataMock: database,
                GetAllBySellerWasCalled: false,
	}

        var resp []Product

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/api/v1/products?seller_id=FEX112AC", "")

	// Act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)

        // fmt.Printf("--- %+v\n", rr)

	// Assert
	assert.Nil(t, err)
        assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, mockService.DataMock[0], resp[0])
}



func TestHandlerGetProductsFail(t *testing.T) {
	// arrange
        errMessage := "cant read database"

        mockService:= MockService{
		DataMock: []Product{},
                Error: errMessage,
	}

	var resp map[string]string
	errorExpected := map[string]string{
		"error": errMessage,
        }	

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/api/v1/products?seller_id=FEX112AC", "")

	// Act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
        
        // fmt.Printf("--- %+v\n", rr)
        // fmt.Printf("--- %+v\n", err)

	// Assert
        assert.Nil(t, err)
        assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, errorExpected, resp)
}

