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
	cols := []string{"Name", "Age", "Email", "Address"}
	rows := sqlmock.NewRows(cols).
		AddRow("test1", 1, "test1@mail.com", "test123").
		AddRow("test2", 2, "test2@mail.com", "test234").
		AddRow("test3", 3, "test3@mail.com", "test345")

	mock.ExpectQuery("SELECT (.+) FROM test.customer").WithArgs("test2@mail.com").WillReturnRows(rows)

	customer := GetSpecificCustomer(db, "test2@mail.com")

	//expected result
	var specificCustomer []models.Customer
	customer1 := models.Customer{
		Name:    "test1",
		Age:     1,
		Email:   "test1@mail.com",
		Address: "test123",
	}
	specificCustomer = append(specificCustomer, customer1)
	customer2 := models.Customer{
		Name:    "test2",
		Age:     2,
		Email:   "test2@mail.com",
		Address: "test234",
	}
	specificCustomer = append(specificCustomer, customer2)
	customer3 := models.Customer{
		Name:    "test3",
		Age:     3,
		Email:   "test3@mail.com",
		Address: "test345",
	}
	specificCustomer = append(specificCustomer, customer3)

	expectedCustomer := models.Customer{
		Name:    "test2",
		Age:     2,
		Email:   "test2@mail.com",
		Address: "test234",
	}
	assert.Equal(t, expectedCustomer, customer)

	defer db.Close()
}
