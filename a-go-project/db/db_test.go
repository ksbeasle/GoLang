package db

import (
	"a-go-project/models"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetSpecificCustomer(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("Error creating mock database:", err)
		return
	}
	defer db.Close()
	cols := []string{"Name", "Age", "Email", "Address"}
	mock.ExpectQuery(`SELECT Name, Age, Email, Address FROM test.customer WHERE Email=?`).
		WithArgs("test2@mail.com").
		WillReturnRows(mock.NewRows(cols).
			AddRow("test2", 2, "test2@mail.com", "test234"))

	customer, _ := GetSpecificCustomer(db, "test2@mail.com")

	expectedCustomer := models.Customer{
		Name:    "test2",
		Age:     2,
		Email:   "test2@mail.com",
		Address: "test234",
	}
	assert.Equal(t, expectedCustomer, customer)
}
