package user_application_service

type RegistorUserCommand struct {
	Name string
}

func NewRegistorUserCommand(name string) (RegistorUserCommand, error) {
	c := RegistorUserCommand{
		Name: name,
	}
	return c, nil
}

type GetUserCommand struct {
	Id string
}

func NewGetUserCommand(id string) (GetUserCommand, error) {
	c := GetUserCommand{
		Id: id,
	}
	return c, nil
}

type UpdateUserCommand struct {
	Id   string
	Name string
}

func NewUpdateUserCommand(id string, name string) (UpdateUserCommand, error) {
	// TODO: nameはoption引数にする必要がある
	c := UpdateUserCommand{
		Id:   id,
		Name: name,
	}
	return c, nil
}

type DeleteUserCommand struct {
	Id string
}

func NewDeleteUserCommand(id string) (DeleteUserCommand, error) {
	c := DeleteUserCommand{
		Id: id,
	}
	return c, nil
}
