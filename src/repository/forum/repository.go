package forum

import (
	"context"
	"errors"
	"time"

	"github.com/cjmarkham/GoBB/config"
	"github.com/cjmarkham/GoBB/src/domain/forum"
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

func (r *Repository) FindByID(ctx context.Context, id uuid.UUID) (*forum.Forum, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var f *forum.Forum

	if err := r.DB.NewSelect().
		Model(&f).
		Where("id=?", id).
		Scan(ctx); err != nil {
		log.Error().Err(err).Msg("could not fetch forum")
		return nil, errors.New("could not fetch forum")
	}

	return f, nil
}

func (r *Repository) FindBySlug(ctx context.Context, slug string) (*forum.Forum, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var f forum.Forum

	if err := r.DB.NewSelect().
		Model(&f).
		Where("slug=?", slug).
		Scan(ctx); err != nil {
		log.Error().Err(err).Msg("could not fetch forum")
		return nil, errors.New("could not fetch forum")
	}

	return &f, nil
}

func (r *Repository) FindWithParents(ctx context.Context) ([]forum.Forum, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var fs []forum.Forum

	if err := r.DB.NewSelect().
		Model(&fs).
		Relation("Children").
		Where("parent_id=?", uuid.Nil).
		Limit(10).
		OrderExpr("created_at ASC").
		Scan(ctx); err != nil {
		log.Error().Err(err).Msg("could not fetch parent forums")
		return nil, errors.New("could not fetch parent forums")
	}

	return fs, nil
}

func (r *Repository) Create(ctx context.Context, dto forum.DTO) (*forum.Forum, error) {
	return nil, nil
}
