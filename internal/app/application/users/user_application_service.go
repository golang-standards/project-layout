package user_application_service

import (
	"fmt"

	"github.com/huroshotoku/golang-project-layout/internal/app/domain/models/users"
)

type IUserApplicationService interface {
	Registor(command RegistorUserCommand) error
	Get(command GetUserCommand) (*UserData, error)
	Update(command UpdateUserCommand) error
	Delete(command DeleteUserCommand) error
}

// 実装側はprivateに
type userApplicationService struct {
	repo    users.IUserRepository
	service users.UserService
}

func NewUserApplicationService(
	repo users.IUserRepository,
	service users.UserService,
) (IUserApplicationService, error) {
	ua := userApplicationService{
		repo:    repo,
		service: service,
	}
	return ua, nil
}

func (s userApplicationService) Registor(command RegistorUserCommand) error {
	userName, err := users.NewUserName(command.Name)
	if err != nil {
		// errorってどこまで伝搬させる？
		return err
	}
	user, err := users.NewUserByName(*userName)
	if err != nil {
		return err
	}
	exist, err := s.service.Exists(*user)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("ユーザ'%s'はすでに存在しています。", command.Name)
	}
	s.repo.Save(*user)
	return nil
}

func (s userApplicationService) Get(command GetUserCommand) (*UserData, error) {
	targetId, err := users.NewUserId(command.Id)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.Find(*targetId)
	if err != nil {
		return nil, err
	}
	userData := NewUserData(*user)
	return userData, nil
}

func (s userApplicationService) Update(command UpdateUserCommand) error {
	targetId, err := users.NewUserId(command.Id)
	if err != nil {
		return err
	}
	user, err := s.repo.Find(*targetId)
	if err != nil {
		return err
	}
	if command.Name != "" {
		// コマンドに関するnilチェックはドメイン知識と別で必要
		newUserName, err := users.NewUserName(command.Name)
		if err != nil {
			return err
		}
		err = user.ChangeName(*newUserName)
		if err != nil {
			return err
		}
		exist, err := s.service.Exists(*user)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("ユーザ'%s'はすでに存在しています。", command.Name)
		}
	}
	s.repo.Save(*user)
	return nil
}

func (s userApplicationService) Delete(command DeleteUserCommand) error {
	targetId, err := users.NewUserId(command.Id)
	if err != nil {
		return err
	}
	user, err := s.repo.Find(*targetId)
	if err != nil {
		// 対象が見つからない場合に退会成功扱いにする方針もあるが
		// ユーザが見つからなかった時の表現をerrにしているので、このケースではできない
		// errの中身見ればできなくはないが、もし上記のような制御をしたいならerrはnilにするべき
		return err
	}
	s.repo.Delete(*user)
	return nil
}
