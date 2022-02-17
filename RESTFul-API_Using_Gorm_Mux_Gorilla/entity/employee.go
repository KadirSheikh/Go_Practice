package entity

//Employee object for REST(CRUD)
type Employee struct {
	ID        int     `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Age       int     `json:"age"`
	JobTitle  string  `json:"jobTitle"`
	Grade     float32 `json:"grade"`
}
