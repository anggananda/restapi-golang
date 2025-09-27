package interfaces

import (
	"context"
	"restapi-golang/models"
)

type UserRepository interface {
	CheckUserByUsername(ctx context.Context, username string) (*models.User, error)
}
