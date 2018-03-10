package user

import (
	"context"

	"google.golang.org/appengine/datastore"
)

type repo struct {
}

// NewDatastoreRepository return a new Datastore repository
func NewDatastoreRepository() Repository {
	return &repo{}
}

func (r *repo) FindAll(ctx context.Context) (*[]User, error) {
	q := datastore.NewQuery("User")

	users := make([]User, 0)
	if _, err := q.GetAll(ctx, &users); err != nil {
		return &users, err
	}
	return &users, nil
}

func (r *repo) Find(ctx context.Context, u *User) (*User, error) {
	userRec := User{}

	key := datastore.NewKey(ctx, "User", u.Username, 0, nil)
	if err := datastore.Get(ctx, key, &userRec); err != nil {
		return &userRec, err
	}
	return &userRec, nil
}

func (r *repo) Save(ctx context.Context, u *User) (*User, error) {
	key := datastore.NewKey(ctx, "User", u.Username, 0, nil)

	_, err := datastore.Put(ctx, key, u)
	if err != nil {
		return u, err
	}
	return u, nil
}
