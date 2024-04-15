package jsonresponse

import (
	"fmt"

	"github.com/ykpythemind/gomvc/models"
)

type UserList struct {
	Users []*User `json:"users"`
}

type User struct {
	// id is not here because it is not necessary for the JSON response
	Name string `json:"name"`
}

func JSONUsers(modelusers []*models.User) (UserList, error) {
	users := make([]*User, len(modelusers))

	for i, modeluser := range modelusers {
		users[i] = &User{
			// 明示的につめつめする
			Name: fmt.Sprintf("name is %s", modeluser.Name),
		}
	}

	return UserList{Users: users}, nil
}
