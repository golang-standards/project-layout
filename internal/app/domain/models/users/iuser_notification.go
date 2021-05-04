package users

type IUserNotification interface {
	Id(id UserId)
	Name(name UserName)
}
