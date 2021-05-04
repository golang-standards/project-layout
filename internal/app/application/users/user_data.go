package user_application_service

import "github.com/huroshotoku/golang-project-layout/internal/app/domain/models/users"

type UserData struct {
	Id   string
	Name string
}

// リポジトリでDataModelにして返すようにしてるけど、これいる？
func NewUserData(user users.User) *UserData {
	// パラメータが変わった時の修正箇所を一箇所に留めるため、ドメインオブジェクトに依存させる.
	ud := UserData{
		Id:   user.Id,
		Name: user.Name,
	}
	// TODO: ここも本来ポインタにしたくないが、nilを扱うためにポインタにしてる（直し方を後で聞く）
	return &ud
}
