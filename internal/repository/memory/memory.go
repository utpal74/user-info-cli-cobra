package memory

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/user-info-cli-tool/pkg/model"
)

// ErrNotFound returns when there is no user found.
var ErrNotFound = errors.New("not found")

// Repository defines a user in memory repository.
type Repository struct {
	data map[string]model.User
	filePath string
}

// New creates a in memory user repository.
func New() *Repository {
	filePath := "users.json"
	data := make(map[string]model.User)
	if file, err := os.Open(filePath); err == nil {
		defer file.Close()
		json.NewDecoder(file).Decode(&data)
	}
	return &Repository{data: data, filePath: filePath}
}

func (repo *Repository) save() error {
	file, err := os.Create(repo.filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(repo.data)
}

// GetUser - fetch the user based on name.
func (r *Repository) GetUser(ctx context.Context, name string) (*model.User, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("invalid user: name can't be empty")
	}
	user, ok := r.data[name]
	if !ok {
		return nil, ErrNotFound
	}
	result := &model.User{
		Name: user.Name, 
		MobileNo: user.MobileNo,
	}
	return result, nil
}

// CreateUser creates a new user
func (repo *Repository) CreateUser(ctx context.Context, name, mobileNo string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("invalid user: name can't be empty")
	}
	mobileNo = strings.TrimSpace(mobileNo)
	if mobileNo == "" {
		return errors.New("invalid mobileNo: mobileNo can't be empty")
	}
	u := model.User{Name: name, MobileNo: mobileNo}
	repo.data[name] = u
	return repo.save()
}
