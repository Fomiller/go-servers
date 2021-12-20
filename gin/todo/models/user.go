package models

type UserService interface {
	CreateUser(firstName string, lastName string, email string, password string)
	GetUser(firstName string)
}

type User struct {
	ID        string `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstName" gorm:"firstName"`
	LastName  string `json:"lastName" gorm:"lastName"`
	Email     string `json:"email" gorm:"email"`
	Password  string `json:"password" gorm:"password"`
}

type CreateUserInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
