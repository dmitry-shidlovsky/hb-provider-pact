package provider

import (
	"github.com/dmitry-shidlovsky/TestPact/model"
)

type UserRepository struct {
	Users map[string]*model.User
}

func (u *UserRepository) TestInit() {
	u.Users = map[string]*model.User{
		"sally": {
			FirstName: "Dmitry",
			LastName:  "Shidlovsky",
			Username:  "dsh",
			Type: 	   "admin",
			ID:        1,
		},
	}
}

func (u *UserRepository) GetUsers() []model.User {
	var response []model.User

	for _, user := range u.Users {
		response = append(response, *user)
	}

	return response
}

func (u *UserRepository) ByUsername(username string) (*model.User, error) {
	if user, ok := u.Users[username]; ok {
		return user, nil
	}
	return nil, model.ErrNotFound
}

func (u *UserRepository) ByID(ID int) (*model.User, error) {
	for _, user := range u.Users {
		if user.ID == ID {
			return user, nil
		}
	}
	return nil, model.ErrNotFound
}
