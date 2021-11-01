package topic

import (
	"context"
	"errors"
	"time"

	"github.com/cjmarkham/GoBB/config"
	"github.com/cjmarkham/GoBB/src/domain/topic"
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
		timeout: time.Duration(config.Timeout) * time.Millisecond,
	}
}

func (r *Repository) FindByForum(ctx context.Context, id uuid.UUID) ([]topic.Topic, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var ts []topic.Topic

	if err := r.DB.NewSelect().
		Model(&ts).
		Relation("Author").
		Relation("LastPoster").
		Where("forum_id=?", id).
		Scan(ctx); err != nil {
		log.Error().Err(err).Msg("could not fetch forum topics")
		return nil, errors.New("could not fetch forum topics")
	}

	return ts, nil
}

func (r *Repository) FindBySlug(ctx context.Context, slug string) (*topic.Topic, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var t topic.Topic

	if err := r.DB.NewSelect().
		Model(&t).
		Where("slug=?", slug).
		Scan(ctx); err != nil {
		log.Error().Err(err).Msg("could not fetch topic")
		return nil, errors.New("could not fetch topic")
	}

	return &t, nil
}
