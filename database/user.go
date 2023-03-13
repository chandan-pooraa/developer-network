// Contains the DB structure of user table

package database

type User struct {
	Id       float64 `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"-"`
	Email    string  `json:"email"`
}
