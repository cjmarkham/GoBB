package forum

import (
	"context"
	"errors"
	"time"

	"github.com/cjmarkham/GoBB/config"
	"github.com/cjmarkham/GoBB/src/domain/forum"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

type Repository struct {
	DB      *bun.DB
	timeout time.Duration
}

func ProvideRepository(config *config.PostgresqlConfig, db *bun.DB) *Repository {
	return &Repository{
		DB:      db,
		timeout: time.Duration(config.Timeout)*time.Millisecond,
	}
}

func (r *Repository) FindAll(ctx context.Context) ([]forum.Forum, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var forums []forum.Forum

	if err := r.DB.NewSelect().Model(&forums).Limit(10).Scan(ctx); err != nil {
		log.Error().Err(err).Msg("could not fetch forums")
		return nil, errors.New("could not fetch forums")
	}

	return forums, nil
}
