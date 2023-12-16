package controllers_test
import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr"
	"book-inventory-system/controllers"
)

// create operation - mock req , 
// create input - validate - data is matching or not 
// http code 
// validate response status

func TestCreateBook_Success(t *testing.T) {
	// Mock context for a successful book creation
	ctx := createMockContext(http.MethodPost, nil, "Title=TestTitle&Author=TestAuthor&Price=10&QuantityAvailable=100", t)

	// Test controllers.CreateBook function
	result, err := controllers.CreateBook(ctx)

	// Assertions
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, "Successfully Created Book", result)
}

func TestCreateBook_InvalidFormData(t *testing.T) {
	// Mock context for invalid form data
	ctx := createMockContext(http.MethodPost, nil, "InvalidFormData", t)

	// Test controllers.CreateBook function
	result, err := controllers.CreateBook(ctx)

	// Assertions
	assert.NotNil(t, err, "Expected an error")
	assert.Contains(t, err.Error(), "invalid URL escape")
	assert.Nil(t, result, "Expected result to be nil")
}

// Utility function to create a mock context for testing
func createMockContext(method string, params map[string]string, body string, t *testing.T) *gofr.Context {
	req := httptest.NewRequest(method, "/dummy", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for key, value := range params {
		req.URL.Query().Add(key, value)
	}

	gofrContext := gofr.NewContext(nil, gofr.NewRequest(req), nil)

	return gofrContext
}

func TestGetAllBooks_Success(t *testing.T) {
	// Mock context for successful GetAllBooks
	ctx := createMockContext(http.MethodGet, nil, "", t)

	// Test controllers.GetAllBooks function
	result, err := controllers.GetAllBooks(ctx)

	// Assertions
	assert.Nil(t, err, "Expected no error")
	assert.NotNil(t, result, "Expected result to be not nil")
	// Add more assertions based on your actual data model
}

func TestGetBook_Success(t *testing.T) {
	// Mock context for successful GetBook
	ctx := createMockContext(http.MethodGet, map[string]string{"Title": "TestTitle"}, "", t)

	// Test controllers.GetBook function
	result, err := controllers.GetBook(ctx)

	// Assertions
	assert.Nil(t, err, "Expected no error")
	assert.NotNil(t, result, "Expected result to be not nil")
	// Add more assertions based on your actual data model
}

func TestUpdateBookQuantity_Success(t *testing.T) {
	// Mock context for successful UpdateBookQuantity
	ctx := createMockContext(http.MethodPatch, nil, "Title=TestTitle&QuantityAvailable=50", t)

	// Test controllers.UpdateBookQuantity function
	result, err := controllers.UpdateBookQuantity(ctx)

	// Assertions
	assert.Nil(t, err, "Expected no error")
	assert.NotNil(t, result, "Expected result to be not nil")
	assert.Equal(t, "Successfully Updated Quantity Of Book", result)
	// Add more assertions based on your actual data model
}

func TestDeleteBook_Success(t *testing.T) {
	// Mock context for successful DeleteBook
	ctx := createMockContext(http.MethodDelete, map[string]string{"id": "1"}, "", t)

	// Test controllers.DeleteBook function
	result, err := controllers.DeleteBook(ctx)

	// Assertions
	assert.Nil(t, err, "Expected no error")
	assert.NotNil(t, result, "Expected result to be not nil")
	assert.Equal(t, "Successfully deleted Book", result)
}