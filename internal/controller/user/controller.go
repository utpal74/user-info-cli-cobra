package user

import (
	"context"
	"errors"

	"github.com/utpal74/user-info-cli-cobra/internal/repository"
	"github.com/utpal74/user-info-cli-cobra/pkg/model"
)

// ErrNotFound returns when there is no user found.
var ErrNotFound = errors.New("not found")

type userRepository interface {
	GetUser(context.Context, string) (*model.User, error)
	CreateUser(context.Context, string,string) error
}

// Controller defines a user controller.
type Controller struct {
	repo userRepository
}

// New creates a user controller.
func New(repo userRepository) *Controller {
	return &Controller{repo}
}

// GetUser - returns user from repository.
func (c *Controller) GetUser(ctx context.Context, name string) (*model.User, error) {
	res, err := c.repo.GetUser(ctx, name)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, nil
}

// CreateUser - creates a new user record.
func (c *Controller) CreateUser(ctx context.Context, name, mobileNo string) error {
	return  c.repo.CreateUser(ctx, name, mobileNo)
}
