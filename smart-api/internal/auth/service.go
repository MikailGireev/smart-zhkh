package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const userFile = "./data/user.json"

func LoadUser() ([]User, error) {
	file, err := os.Open(userFile)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func SaveUsers(users []User) error {
	file, err := os.Create(userFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(users)
}

func RegisterUser(user_id, username, password string) error {
	users, err := LoadUser()

	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Username == username {
			return nil
		}
	}

	newUser := User{
		ID:       user_id,
		Username: username,
		Password: password,
	}

	if err := newUser.Validate(); err != nil {
		return err
	}

	users = append(users, newUser)
	return SaveUsers(users)
}

func LoginUser(username, password string) (*User, error) {
	users, err := LoadUser()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}

	return nil, errors.New("invalid credentials")
}

func (u *User) Validate() error {
	if len(u.Username) < 2 {
	fmt.Println(2)
		return fmt.Errorf("username must be at least 2 characters: %w", ErrValidation)
	}
	// if  len(u.Password) < 4 {
	// fmt.Println(3)
	// 	return fmt.Errorf("password must be at least 6 characters: %w", ErrValidation)
	// }
	if u.Username == "" {
	fmt.Println(4)
		return fmt.Errorf("invalid username: %w", ErrValidation)
	}
	if u.Password == "" {
	fmt.Println(5)
		return fmt.Errorf("invalid password: %w", ErrValidation)
	}
	return nil
}