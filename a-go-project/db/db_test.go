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

	/* NOTE: TEST FAILS USING THIS WAY FOR SOME REASON, IT ALWAYS RETURNS THE FIRST ENTRY AND
	NEVER REALLY LOOKS FOR THAT SECOND ENTRY, DON'T KNOW WHY.
	*/
	// rows := sqlmock.NewRows(cols).
	// 	AddRow("test1", 1, "test1@mail.com", "test123").
	// 	AddRow("test2", 2, "test2@mail.com", "test234").
	// 	AddRow("test3", 3, "test3@mail.com", "test345")
	// mock.ExpectQuery("^SELECT (.+) FROM test.customers*").WithArgs("test2@mail.com").WillReturnRows(rows)

	//This works though.
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
