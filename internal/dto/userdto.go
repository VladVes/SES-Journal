package dto

import (
	"errors"
	"net/mail"
	"strings"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Active   bool   `json:"active"`
}

func (in *CreateUserRequest) Normalize() {
	in.Name = strings.TrimSpace(in.Name)
	in.Name = strings.TrimSpace(in.Password)
	in.Name = strings.TrimSpace(in.Role)
	in.Email = strings.ToLower(strings.TrimSpace(in.Email))
}

func (in *CreateUserRequest) Validate() error {
	if in.Name == "" {
		return errors.New("name is required")
	}
	if in.Email == "" {
		return errors.New("email is required")
	}
	if in.Password == "" {
		return errors.New("password is required")
	}

	addr, err := mail.ParseAddress(in.Email)
	if err != nil {
		return errors.New("invalid email format")
	}
	//  запретить "Name <email@...>" форму
	if addr.Address != in.Email {
		return errors.New("invalid email format")
	}

	return nil
}

type CreateUserResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Active bool   `json:"active"`
}
