package post

import (
	"context"
	"errors"
	"time"

	"github.com/cjmarkham/GoBB/config"
	"github.com/cjmarkham/GoBB/src/domain/post"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

type Repository struct {
	DB      *bun.DB
	timeout time.Duration
}

type Model struct {
	ID string
}

func ProvideRepository(config *config.PostgresqlConfig, db *bun.DB) *Repository {
	return &Repository{
		DB:      db,
		timeout: time.Duration(config.Timeout)*time.Millisecond,
	}
}

func (r *Repository) FindByTopic(ctx context.Context, id uuid.UUID) ([]post.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var ps []post.Post

	if err := r.DB.NewSelect().
		Model(&ps).
		Relation("Author").
		Where("topic_id=?", id).
		Scan(ctx);
		err != nil {
		log.Error().Err(err).Msg("could not fetch topic posts")
		return nil, errors.New("could not fetch topic posts")
	}

	return ps, nil
}
