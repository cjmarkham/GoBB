package forum

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Forum, error)
}
