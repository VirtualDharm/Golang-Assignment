// user_microservice/user.go
package user_microservice

import (
	"errors"
	"fmt"
)

type UserType string

const (
	DefaultUser UserType = "default"
	AdminUser   UserType = "admin"
)

type User struct {
	Username string
	Email    string
	Type     UserType
}

type UserMicroservice struct {
	Users []User
}

func NewUserMicroservice() *UserMicroservice {
	return &UserMicroservice{}
}

func (um *UserMicroservice) AddUser(adminUser User, username, email string, userType UserType) error {
	if adminUser.Type != AdminUser {
		return errors.New("only admin can add users")
	}

	user := User{
		Username: username,
		Email:    email,
		Type:     userType,
	}
	um.Users = append(um.Users, user)
	fmt.Println("User added:", user)
	return nil
}

func (um *UserMicroservice) DeleteUser(adminUser User, username string) error {
	if adminUser.Type != AdminUser {
		return errors.New("only admin can delete users")
	}

	for i, user := range um.Users {
		if user.Username == username {
			um.Users = append(um.Users[:i], um.Users[i+1:]...)
			fmt.Println("User deleted:", username)
			return nil
		}
	}
	return errors.New("user not found")
}

func (um *UserMicroservice) GetUser(currentUser User, username string) (*User, error) {
	if currentUser.Type != AdminUser && currentUser.Username != username {
		return nil, errors.New("unauthorized to fetch this user's information")
	}

	for _, user := range um.Users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (um *UserMicroservice) GetAllUsers(adminUser User) ([]User, error) {
    if adminUser.Type != AdminUser {
        return nil, errors.New("only admin can get all users")
    }
    return um.Users, nil
}

func (u *User) IsAdmin() bool {
	return u.Type == AdminUser
}
