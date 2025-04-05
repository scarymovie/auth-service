package repository

import (
	userpb "auth-service/proto"
)

type UserRepository interface {
	Create(name string) (*userpb.User, error)
	FindById(id string) (*userpb.User, error)
}
