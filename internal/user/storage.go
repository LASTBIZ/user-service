package user

type Storage interface {
	CreateUser(user User) (User, error)
	DeleteUser(userId uint32) error
	UpdateUser(user User) (User, error)
	GetUser(userId uint32) (User, error)
	GetUserByEmail(email string) (User, error)
}
