package configs_test

import (
	"go-crud/configs"
	"go-crud/tests/mocks"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {

	t.Run("Verify Database connection", func(t *testing.T) {
		// Create a new instance of the mock DB connection
		mockDB := new(mocks.MockDBConnection)

		// Set the expected behavior of the Connection method
		expectedDB := &gorm.DB{}

		mockDB.On("Connection").Return(expectedDB)

		// Replace the original implementation with the mock implementation
		configs.Connection = mockDB.Connection

		// Call the Connection function
		db := configs.Connection()

		// Ensure that the returned DB object matches the expected DB object
		assert.Equal(t, expectedDB, db)

		// Call the Connection method on the mock DB
		mockDB.AssertExpectations(t)
	})
}
