package models

// User : defines a user from the ccuser table
type User struct {
	Id        int
	Login     string
	FirstName string
	LastName  string
	Email     string
}
