package user

import (
	"context"
)

type testRepo struct {
}

// NewTestRepository return a new test repository
func NewTestRepository() Repository {
	return &testRepo{}
}

func (r *testRepo) FindAll(ctx context.Context) (*[]User, error) {
	users := make([]User, 0)
	return &users, nil
}

func (r *testRepo) Find(ctx context.Context, u *User) (*User, error) {
	newUser := User{}

	println("SER")
	println(u.BonusSelectionPoints)

	// copy values of u into newUser.
	// TODO: this function should not take a pointer to
	// u; just pass in a copy (u User), not (u *User)
	// to avoid mutating the passed-in struct, which messes
	// up tests.
	newUser = *u
	hashed, _ := HashPassword(u.Password)
	newUser.Password = hashed
	return &newUser, nil
}

func (r *testRepo) Save(ctx context.Context, u *User) (*User, error) {
	return u, nil
}
