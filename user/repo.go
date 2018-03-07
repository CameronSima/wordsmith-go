package user

import "context"

// Repository is the interface for User DAO
type Repository interface {
	FindAll(ctx context.Context) (*[]User, error)
	Find(ctx context.Context, user *User) (*User, error)
	Save(ctx context.Context, user *User) (*User, error)
}
