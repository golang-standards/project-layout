package users

import "fmt"

type UserId string

func NewUserId(value string) (*UserId, error) {
	// ここにルールを指定することで、不正なデータを作成できないようにする
	if value == "" {
		return nil, fmt.Errorf("ValueError: UserIdがnullです。")
	}
	id := UserId(value)
	return &id, nil
}
