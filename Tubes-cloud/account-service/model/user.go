package model

//Representasi data user
//Dipakai di semua layer (repository, service)
type User struct {
	ID       int
	Email    string
	Password string
}