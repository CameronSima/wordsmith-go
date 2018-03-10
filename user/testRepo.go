package user

import (
	"context"
)

type testRepo struct {
}

// NewTestRepository return a new Datastore repository
func NewTestRepository() Repository {
	return &testRepo{}
}

func (r *testRepo) FindAll(ctx context.Context) (*[]User, error) {
	users := make([]User, 0)
	return &users, nil
}

func (r *testRepo) Find(ctx context.Context, u *User) (*User, error) {
	newUser := User{}
	hashed, _ := HashPassword(u.Password)
	newUser.Password = hashed
	return &newUser, nil
}

func (r *testRepo) Save(ctx context.Context, u *User) (*User, error) {
	return u, nil
}
