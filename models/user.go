package models

type User struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
	Password string `gorm:"not null;default:null"`
}

// in our userService we will need the following 3 methods
// GetUserById | Register | Login

// and in the userRepository we will have:
// GetUserById | CreateUser | FindUserByEmail

type IUserRepository interface {
	GetUserById(id int) (*User, error)
	CreateUser(user *User) (*User, error)
	FindUserByEmail(email string) (*User, error)
}

type IUserService interface {
	GetUserById(id int) (*User, error)
	Register(user *User) (*User, error)
	Login(email, password string) (*User, error)
}