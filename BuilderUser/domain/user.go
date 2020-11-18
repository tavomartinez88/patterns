//domain is a package that contains all types representant a object real life
package domain

// User is a type that represent a Person
type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone"`
}
