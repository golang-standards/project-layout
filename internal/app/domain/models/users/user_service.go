package users

type UserService struct {
	userRepository IUserRepository
}

func NewUserSirvice(userRepository IUserRepository) (UserService, error) {
	us := UserService{userRepository: userRepository}
	return us, nil
}

func (us UserService) Exists(user User) (bool, error) {
	found, err := us.userRepository.FindByName(user.name)
	if err != nil {
		return false, err
	}
	return found != nil, err
}
