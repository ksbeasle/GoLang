package models

//Customer struct
type Customer struct {
	name 	string 	`json:"name"`
	age 	int		`json:"age"`
	email 	string 	`json:"email"`
	address string	`json:"address"`
}