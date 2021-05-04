package users

type IUserRepository interface {
	Find(id UserId) (*User, error)
	FindByName(name UserName) (*User, error)
	Save(user User) error
	Delete(user User) error
}
