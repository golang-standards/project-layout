package users

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

type User struct {
	id   UserId
	name UserName
}

func NewUser(id UserId, name UserName) (*User, error) {
	if id == "" {
		return nil, fmt.Errorf("ValueError: UserIdがnullです。")
	}
	if name == "" {
		return nil, fmt.Errorf("ValueError: UserNameがnullです。")
	}
	user := User{
		id:   id,
		name: name,
	}
	return &user, nil
}

func NewUserByName(name UserName) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("ValueError: UserNameがnullです。")
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("UserIdの生成に失敗しました: %s", err)
	}
	user := User{
		id:   UserId(id.String()),
		name: name,
	}
	return &user, nil
}

func (u *User) ChangeName(name UserName) error {
	if name == "" {
		return fmt.Errorf("ValueError: UserNameがnullです。")
	}
	u.name = name
	return nil
}

func (u User) Equals(other User) (bool, error) {
	if (other == nil) || reflect.ValueOf(other).IsNil() {
		return false, nil
	}
	return u.id == other.id, nil
}

func (u User) Notify(note IUserNotification) {
	note.Id(u.id)
	note.Name(u.name)
}
