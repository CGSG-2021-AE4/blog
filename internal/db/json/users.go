package json

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"webapp/api"

	"github.com/google/uuid"
)

type UserStore struct {
	filename string

	mutex sync.Mutex
	users map[uuid.UUID]*api.User
}

func NewUserStore(filename string) (*UserStore, error) {
	us := UserStore{filename: filename, users: make(map[uuid.UUID]*api.User)}
	if err := us.load(); err != nil {
		return nil, err
	}
	return &us, nil
}

func (us *UserStore) load() error {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	bytes, err := os.ReadFile(us.filename)
	if err != nil {
		return err
	}
	users := []api.User{}
	if err := json.Unmarshal(bytes, &users); err != nil {
		return nil
	}
	for _, u := range users { // Check if this is memory friendly
		us.users[u.Id] = &u
	}
	return nil
}

func (us *UserStore) save() error {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	users := []api.User{}
	for _, u := range us.users {
		users = append(users, *u)
	}
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return os.WriteFile(us.filename, bytes, 0777)
}

func (us *UserStore) GetUser(ctx context.Context, id uuid.UUID) (*api.User, error) {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	u := us.users[id]
	if u == nil {
		return nil, api.ErrUserNotExists
	}
	return u, nil
}

func (us *UserStore) DoExist(ctx context.Context, username string) (bool, error) {
	// In this implementation I can just do this
	if _, err := us.GetUserByName(ctx, username); err != nil {
		return false, nil
	}
	return true, nil
}

func (us *UserStore) GetUserByName(ctx context.Context, username string) (*api.User, error) {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	for id := range us.users { // If I understand right this way I will not copy map
		if us.users[id].Username == username {
			return us.users[id], nil
		}
	}
	return nil, api.ErrUserNotExists
}

func (us *UserStore) CreateUser(ctx context.Context, user *api.User) error {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	// Assume that username is unique
	us.users[user.Id] = user
	return nil
}

func (us *UserStore) DeleteUser(ctx context.Context, Id uuid.UUID) error {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	if us.users[Id] == nil {
		return api.ErrUserNotExists
	}
	delete(us.users, Id)
	return nil
}

func (us *UserStore) Close() error {
	return us.save()
}
