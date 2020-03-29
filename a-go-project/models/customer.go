package models

//Customer struct
type Customer struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
